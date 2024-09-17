package cmd

import (
	"fmt"
	"ldapctl/g"
	"ldapctl/models"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var csvFile bool

func init() {
	searchMultiCmd.Flags().BoolVarP(&csvFile, "file", "f", false, "output search to users.csv, failed search to failed.csv")
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchFilterCmd)
	searchCmd.AddCommand(searchUserCmd)
	searchCmd.AddCommand(searchMultiCmd)
}

var searchCmd = &cobra.Command{
	Use:       "search",
	Short:     "Search Test",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{searchFilterCmd.Use, searchUserCmd.Use, searchMultiCmd.Use},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
			filter      Search By Filter
			multi       Search Multi Users
			user        Search Single User
			`)
	},
}

var searchFilterCmd = &cobra.Command{
	Use:   "filter [searchFilter]",
	Short: "Search By Filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search By Filter"
		filter := args[0]
		startTime := time.Now()
		PrintStart(action)

		results, err := models.SingleSearch(g.Config().Ldap, filter)
		if err != nil {
			fmt.Printf("%s Search failed: %s \n", filter, err.Error())
			PrintEnd(action, startTime)
			return
		}
		for _, result := range results {
			PrintSearchResult(result)
		}
		fmt.Println("results count ", len(results))
		PrintEnd(action, startTime)
	},
}

var searchUserCmd = &cobra.Command{
	Use:   "user [username]",
	Short: "Search Single User",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search Single User"

		username := args[0]
		startTime := time.Now()
		PrintStart(action)

		result, err := models.SingleSearchUser(g.Config().Ldap, username)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s Search failed: %s \n", username, err.Error())
			PrintEnd(action, startTime)
			return
		}
		PrintSearchResult(result)

		PrintEnd(action, startTime)
	},
}

var searchMultiCmd = &cobra.Command{
	Use:   "multi [filename]",
	Short: "Search Multi Users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Multi Search Users"

		userList := args[0]
		searchUsers, err := g.GetLines(userList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read file %s failed: %s \n", userList, err.Error())
			return
		}
		startTime := time.Now()
		PrintStart(action)

		res, err := models.MultiSearchUser(g.Config().Ldap, searchUsers)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Multi Search failed: %s \n", err.Error())
			PrintEnd(action, startTime)
			return
		}
		if csvFile {
			err = WriteUsersToCsv(res.Users, g.USERS_CSV)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Open file %s failed: \n", g.USERS_CSV)
				return
			}
			if len(res.FailedMessages) > 0 {
				err = WriteFailsToCsv(res.FailedMessages, g.FAILED_CSV)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Open file %s failed: \n", g.FAILED_CSV)
					return
				}
			}
			fmt.Println("OutPut to csv successed")
			PrintEnd(action, startTime)
			return
		}
		fmt.Println("Successed users:")
		for _, user := range res.Users {
			PrintSearchResult(user)
		}

		for _, failedMessage := range res.FailedMessages {
			fmt.Printf("%s : %s \n", failedMessage.Username, failedMessage.Message)
		}
		fmt.Println()
		fmt.Printf("Successed count %d \n", res.Successed)
		fmt.Printf("Failed count %d \n", res.Failed)

		PrintEnd(action, startTime)
	},
}

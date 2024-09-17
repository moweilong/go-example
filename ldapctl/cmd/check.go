package cmd

import (
	"fmt"
	"ldapctl/g"
	"ldapctl/models"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check Ldap Connectivity",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := models.HealthCheck(g.Config().Ldap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
			return
		}
		fmt.Println("Successed")
	},
}

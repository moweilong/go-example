package cmd

import (
	"fmt"
	"ldapctl/g"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ldapctl",
	Long:  `All software has versions. This is ldapctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ldapctl version : ", g.VERSION)
	},
}

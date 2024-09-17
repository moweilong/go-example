package cmd

import (
	"ldapctl/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Enable a http server for ldapctl",
	Long:  `ldapctl provide a restful api for ldap test`,
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}

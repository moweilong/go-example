package cmd

import (
	"fmt"
	"ldapctl/g"
	"os"

	"github.com/spf13/cobra"
)

var cfg string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfg, "config", "c", "conf.json", "load config file. default conf.json")
}

func initConfig() {
	g.ParseConfig(cfg)
}

var rootCmd = &cobra.Command{
	Use:   "ldapctl",
	Short: "A command line tool for LDAP.",
	Long:  `A command line tool for LDAP.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
  auth        Auth Test
  check	      Check Cdap Connectivity
  help        Help about any command
  http        Enable a http server for ldap-test-tool
  search      Search Test
  version     Print the version number of ldap-test-tool
		`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

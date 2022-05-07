/*
Package cmd
Copyright © 2022 Marc Bouvier <m.bouvier.dev@gmail.com>
    Primary port: Command Line Interface entrypoint.
*/
package cmd

import (
	"baas/contact"
	"baas/talks"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "baas",
	Short: "CLI to get interact with Marc Bouvier's professional informations",
	Long: `BaaS is a CLI tool to interact with up-to-date professional informations about Marc Bouvier.
Such informations are : resume, contact infos, news, curated lists, tip of the day...

Usage:  
  baas [command]

Available Commands:

Flags:
  -h, --help             help for baas

Use "baas [command] --help" for more information about a command.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(contact.Cmd)
	rootCmd.AddCommand(talks.Cmd)
}

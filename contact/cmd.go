/*
Package contact
Copyright © 2022 Marc Bouvier <m.bouvier.dev@gmail.com>
*/
package contact

import (
	"baas/contact/show"
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd represents the contact command
var Cmd = &cobra.Command{
	Use:   "contact",
	Short: "Shows Marc Bouvier's contact infos",
	Long: `Show Marc Bouvier's contact infos:

- Short contact description
- e-mail
- phone
- Web Site
- Social Networks
`,
	Run: func(cmd *cobra.Command, args []string) {
		do, _ := show.Do()
		fmt.Println(do)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

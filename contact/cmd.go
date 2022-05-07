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

func init() {}

/*
Package talks
Copyright © 2022 Marc Bouvier <m.bouvier.dev@gmail.com>
*/
package talks

import (
	"baas/talks/list"
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd represents the talks command
var Cmd = &cobra.Command{
	Use:   "talks",
	Short: "Shows where Marc Bouvier's is speaking in events",
	Long: `Shows where Marc Bouvier's is speaking in events:

- Talk title
    - Event name
    - Event date and hour
    - Event location
    - Event URL
`,
	Run: func(cmd *cobra.Command, args []string) {
		do, _ := list.Do()
		fmt.Println(do)
	},
}

func init() {}

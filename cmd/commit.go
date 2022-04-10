/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/spf13/cobra"
)

var msg *string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		git.Commit(*msg)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	msg = commitCmd.Flags().StringP("message", "m", "", "Commit message")
}

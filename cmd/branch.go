/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a new branch or list existing branches",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			git.Branch("")
		} else if len(args) > 1 {
			fmt.Println("Too many arguments")
		} else {
			git.Branch(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}

/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/spf13/cobra"
)

var bare *bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the current directory as a new repository",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		git.Init(*bare)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	bare = initCmd.Flags().BoolP("bare", "b", false, "Initialize a bare repository")
}

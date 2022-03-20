/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"github.com/mattcarp12/go-gitm/gitm"
	"github.com/spf13/cobra"
)

var bare *bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the current directory as a new repository",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitm.Git{}.Init(*bare)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bare = initCmd.Flags().BoolP("bare", "b", false, "Initialize a bare repository")

}

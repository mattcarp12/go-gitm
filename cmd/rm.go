/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"log"

	"github.com/mattcarp12/go-gitm/gitm"
	"github.com/spf13/cobra"
)

var recurse *bool

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a file from the index and working tree",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No files specified")
		} else if len(args) > 1 {
			log.Fatal("Only specify a single pathSpec")
		} else {
			gitm.Git{}.Rm(args[0], *recurse)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	recurse = rmCmd.Flags().BoolP("recurse", "r", false, "Remove all files in specified directory")
}

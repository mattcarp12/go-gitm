/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package cmd

import (
	"github.com/mattcarp12/go-gitm/gitm/git"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		git.Checkout(args[0])
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

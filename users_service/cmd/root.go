package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "hermes is a cli tool for operating the hermes users service",
	Long:  "hermes is a cli tool for operating the hermes users service i.e creating, modifying and deleting users.",
	Run: func(cmd *cobra.Command, args []string) {
		// This command will create a new user

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Hermes '%s'\n", err)
		os.Exit(1)
	}
}

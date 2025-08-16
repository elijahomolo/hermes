package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Delete a user from the Hermes users service",
	Long:  "Usage: hermes delete-user <user_id>",
	Args:  cobra.ExactArgs(1), // Expect exactly 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		userID := args[0]

		err := userService.DeleteUser(userID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting user: %v\n", err)
			return
		}

		fmt.Printf("User with ID %s deleted successfully\n", userID)
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}

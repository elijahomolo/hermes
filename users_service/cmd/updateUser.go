package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Update a user's information in the Hermes users service",
	Long:  "Usage: hermes update-user <user_id>",
	Args:  cobra.ExactArgs(8), // Expect exactly 8 arguments
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 8 {
			fmt.Println("All fields are required: user_id, first_name, last_name, date_of_birth, password, country, language, email")
			return
		}

		if len(args) > 8 {
			fmt.Println("Too many arguments. Expected 8 arguments.")
			return
		}

		_, err := userService.UpdateUser(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating user: %v\n", err)
			return
		}

		fmt.Printf("User with ID %s updated successfully\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(updateUserCmd)
}

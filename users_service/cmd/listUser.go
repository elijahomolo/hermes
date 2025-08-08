package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listAllUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "List all users in the Hermes users service",
	Long:  "This command lists all users stored in the Hermes users service.",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := userService.ListAllUsers()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing users: %v\n", err)
			return
		}
		if len(users) == 0 {
			fmt.Println("No users found.")
			return
		}
		fmt.Println("List of users:")
		for _, user := range users {
			fmt.Printf("ID: %s, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
		}

	},
}

func init() {
	rootCmd.AddCommand(listAllUsersCmd)
}

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/elijahomolo/hermes/users_service/users"
	"github.com/spf13/cobra"
)

var userService *users.Service

var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "hermes is a cli tool for creating users in the hermes users service",
	Long:  "Usage: hermes-create-user <first_name> <last_name> <date_of_birth> <password> <country> <language> <email>",
	Args:  cobra.ExactArgs(7), // Expect exactly 7 arguments
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 7 {
			fmt.Println("All fields are required: first_name, last_name, date_of_birth, password, country, language, email")
			return
		}

		if len(args) > 7 {
			fmt.Println("Too many arguments. Expected 7 arguments.")
			return
		}

		createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")

		// Call the CreateUser function from the service

		user, err := userService.CreateUser(args[0], args[1], args[2], args[3], args[4], args[5], args[6], createdAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating user: %v\n", err)
			return
		}

		fmt.Printf("User created: %+v\n", user)
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}

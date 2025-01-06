package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	models "github.com/unnxt30/LM-Notif/internal/model"
	memory "github.com/unnxt30/LM-Notif/internal/store"
)

func ValidateAdminRole(user *models.User) error {
	if user.Role != models.AdminRole {
		return errors.New("user does not have admin role")
	}
	return nil
}

func NewUser(name string, role models.Role)	(*models.User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if role != models.AdminRole && role != models.UserRole {
		return nil, errors.New("invalid role")
	}


	return &models.User{
		Name: name,
		Role: role,
		SubTopics: make(map[string]bool),
	}, nil
}

var addUserCmd = &cobra.Command{
	Use: "addUser [userName] [role]",
	Short: "Add a user to the system",
	Long: `Add a user to the system`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		userName := args[0]
		role := args[1]
	
		// Convert role string to models.Role type
		var userRole models.Role
		switch role {
		case "ADMIN":
			userRole = models.AdminRole
		case "USER":
			userRole = models.UserRole
		default:
			return errors.New("invalid role")
		}

		// Create a new user
		user, err := NewUser(userName, userRole)
		if err != nil {
			return err
		}

		err = memory.GlobalMemoryStore.AddUser(user)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("User Succesfully added")
		return nil
	},
}

var getUsersCmd = &cobra.Command{
	Use: "getUsers",
	Short: "Get all users in the system",
	Long: `Get all users in the system`,
	Run: func(cmd *cobra.Command, args []string) {
		users := memory.GlobalMemoryStore.GetAllUsers()
		for _, user := range users {
			fmt.Printf("%v - %v\n", user.Name, user.Role)
		}
	},
}

var viewSubscribedTopicscmd = &cobra.Command{
	Use: "viewSubscribedTopics [userName]",
	Short: "View all topics a user is subscribed to",
	Long: `View all topics a user is subscribed to`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userName := args[0]
		user, err := memory.GlobalMemoryStore.GetUser(userName)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		for topic := range user.SubTopics {
			fmt.Println(topic)
		}
		return nil
	},
}

var removeUserCmd = &cobra.Command{
	Use: "removeUser [userName] [byUserName]",
	Short: "Remove a user from the system",
	Long: `Remove a user from the system`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		userName := args[0]
		byUserName := args[1]

		_, err := memory.GlobalMemoryStore.GetUser(userName)
		if err != nil {
			fmt.Println(err)
			return nil
		}


		byUser, err := memory.GlobalMemoryStore.GetUser(byUserName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if byUser.Name == userName {
			fmt.Println("You cannot remove yourself")
			return nil
		}

		if err := ValidateAdminRole(byUser); err != nil {
			fmt.Println(err)
			return nil
		}
		
		err = memory.GlobalMemoryStore.RemoveUser(userName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println("User successfully removed")
		return nil
	},
}

func init(){
	rootCmd.AddCommand(addUserCmd)
	rootCmd.AddCommand(getUsersCmd)
	rootCmd.AddCommand(viewSubscribedTopicscmd)
	rootCmd.AddCommand(removeUserCmd)
}
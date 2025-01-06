package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	models "github.com/unnxt30/LM-Notif/internal/model"
	memory "github.com/unnxt30/LM-Notif/internal/store"
)

func NewTopic(topicName string) (*models.Topic, error) {
	if topicName == "" {
		return nil, errors.New("topic name cannot be empty")
	}

	return &models.Topic{
		TopicName: topicName,
		UsersSubscribed: make([]models.User, 0),
	}, nil
}

var addTopicCmd = &cobra.Command{
	Use: "addTopic [topicName] [caller]",
	Short: "Add a topic to the system",
	Long: `Add a topic to the system`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		topicName := args[0]
		caller := args[1]
		
		user, err := memory.GlobalMemoryStore.GetUser(caller)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = ValidateAdminRole(user)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Create a new topic
		topic, err := NewTopic(topicName)
		if err != nil {
			return err
		}
		// Add the topic to the store
		if err := memory.GlobalMemoryStore.AddTopic(topic); err != nil {
			return err
		}
		fmt.Println("topic successfully added")
		return nil
	},
}


var getTopicsCmd = &cobra.Command{
	Use: "getTopics",
	Short: "Get all topics in the system",
	Long: `Get all topics in the system`,
	Run: func(cmd *cobra.Command, args []string) {
		topics := memory.GlobalMemoryStore.GetAllTopics()
		for _, topic := range topics {
			fmt.Println(topic.TopicName)
		}
	},
}

var subScribeTopicCmd = &cobra.Command{
	Use: "subscribeTopic [topicName] [userName]",
	Short: "Subscribe a user to a topic",
	Long: `Subscribe a user to a topic`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		topicName := args[0]
		userName := args[1]

		user, err := memory.GlobalMemoryStore.GetUser(userName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		_, err = memory.GlobalMemoryStore.GetTopic(topicName)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		
		memory.GlobalMemoryStore.AddUserToTopic(topicName, user)
		user.SubTopics[topicName] = true
		fmt.Println("topic successfully subscribed")
		return nil
	},
}

var removeTopicCmd = &cobra.Command{
	Use: "removeTopic [topicName] [caller]",
	Short: "Remove a topic from the system",
	Long: `Remove a topic from the system`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		topicName := args[0]
		caller := args[1]
		
		user, err := memory.GlobalMemoryStore.GetUser(caller)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = ValidateAdminRole(user)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		_, err = memory.GlobalMemoryStore.GetTopic(topicName)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		
		err = memory.GlobalMemoryStore.RemoveTopic(topicName)		

		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println("topic successfully removed")
		return nil
	},
}


func init() {
	rootCmd.AddCommand(addTopicCmd)
	rootCmd.AddCommand(getTopicsCmd)
	rootCmd.AddCommand(subScribeTopicCmd)
}
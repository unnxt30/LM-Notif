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

		return nil
	},
}
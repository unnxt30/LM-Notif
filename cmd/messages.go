package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	models "github.com/unnxt30/LM-Notif/internal/model"
	memory "github.com/unnxt30/LM-Notif/internal/store"
)


var publishMessageCmd = &cobra.Command{
	Use: "publishMessage [json]",
	Short: "Publish a message to a topic",
	Long: `Publish a message to a topic`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		obj := args[0]
		var message models.Message	
		if err := json.Unmarshal([]byte(obj), &message); err != nil {
			return fmt.Errorf("error unmarshalling message: %v", err) 
		}

		topic, err := memory.GlobalMemoryStore.GetTopic(message.TopicName)
		if err != nil {
			return fmt.Errorf("error getting topic: %v", err)
		}

		for _, user := range topic.UsersSubscribed {
			fmt.Printf("{\n \t\"topic\": \"%s\",\n \t\"message\": \"%s\",\n \t\"sentTo\": \"%s\"\n}\n", topic.TopicName, message.Text, user.Name) 
		}


		return nil
	},
}

func init(){
	rootCmd.AddCommand(publishMessageCmd)
}
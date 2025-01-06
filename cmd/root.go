package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "notification-service",
	Short: "A notification service CLI Application for Publication notificaitons",
	Long: `A notification service CLI Application for Publication notificaitons`,
}

func init() {
	rootCmd.Run = runRootCmd
}

func runRootCmd(cmd *cobra.Command, args []string) {
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("> ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	if input == "quit" || input == "exit" {
	// 		fmt.Println("Exiting...")
	// 		break
	// 	}
	// 	rootCmd.SetArgs(strings.Split(input, " "))
	// 	if err := rootCmd.Execute(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	reader := bufio.NewReader(os.Stdin)
        for {
            fmt.Print("> ")
            input, _ := reader.ReadString('\n')
            input = strings.TrimSpace(input)
            
            if input == "quit" || input == "exit" {
                fmt.Println("Exiting...")
                break
            }
            
            // Split preserving quoted strings
            r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
            parts := r.FindAllString(input, -1)
            
            // Clean up quotes
            for i, part := range parts {
                parts[i] = strings.Trim(part, "'\"")
            }
            
            rootCmd.SetArgs(parts)
            if err := rootCmd.Execute(); err != nil {
                fmt.Println("Error:", err)
            }
        }
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
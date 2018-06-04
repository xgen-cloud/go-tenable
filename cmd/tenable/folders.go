package tenable

import (
	"context"
	"fmt"
	"log"

	tenableClient "github.com/mistsys/go-tenable/client"
	"github.com/spf13/cobra"
)

var foldersCmd = &cobra.Command{
	Use:   "folders COMMAND",
	Short: "Use the Tenable folders API",
	Args:  cobra.MinimumNArgs(1),
}

// fooCmd represents the foo command
var foldersListCmd = &cobra.Command{
	Use:   "list [ID...]",
	Short: "List folders.",
	Run: func(cmd *cobra.Command, args []string) {
		client = tenableClient.NewClient(accessKey, secretKey)
		client.Debug = debug
		lst, err := client.FoldersList(context.Background())
		if err != nil {
			log.Println("Error getting folders list", err)
		}
		fmt.Printf("%q", lst)
	},
}

func init() {
	rootCmd.AddCommand(foldersCmd)
	foldersCmd.AddCommand(foldersListCmd)
}

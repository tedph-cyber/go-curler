package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var curleCmd = &cobra.Command{
	Use:     "Copy content off the web into local file on local machine",
	Aliases: []string{"curl"},
	Short:   "download files",
	Long:    "Copy content off webpages and paste them in a locally created file",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No URL provided")
			os.Exit(1)
		}
		link := args[0]
		body, err := Copy(link)
		if err != nil {
			log.Fatal(err)
		}
		filename := args[1]
		fmt.Println("Curling of webpage...\n", args[0], args[1], Paste(body, filename))
	},
}

func init() {
	rootCmd.AddCommand(curleCmd)
}

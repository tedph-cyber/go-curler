package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var curleCmd = &cobra.Command{
	Use:     "Copy content off the web into local file on local machine",
	Aliases: []string{"curl"},
	Short:   "download files",
	Long:    "Copy content off webpages and paste them in a locally created file",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
		if len(args) == 0 {
			fmt.Println("Error: No URL provided")
			os.Exit(1)
		}
		s.Start()
		s.Suffix = color.YellowString(" Curling of webpage.\n")
		link := args[0]
		body, err := Copy(link)
		if err != nil {
			s.Stop()
			log.Fatal(err)
		}
		filename := args[1]		
		s.Stop()
		fmt.Println("File from ", args[0], "is downloaded. \n", Paste(body, filename))
		fmt.Println("Your file is now available at " + color.GreenString(filename))
	},
}

func init() {
	rootCmd.AddCommand(curleCmd)
}

package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "curler",
    Short: "Curler is a cli tool for download of files",
    Long:  "Curler is a cli tool used for download of files from the internet by pasting the download URL and it automatically cllones to the local machine withthe same name",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing Curler '%s'\n", err)
        os.Exit(1)
    }
}

package cmd

import "github.com/spf13/cobra"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize db and table",
	Long:  "Initialize a new studybuddy database and table",
}

func init() {
	rootCmd.AddCommand(initCmd)

}

package cmd

import (
	"github.com/spf13/cobra"
)

// var rootCmd = &cobra.Command{
// 	Use:   "studybuddy",
// 	Short: "this is a short version of doing the commmand ayo my brother",
// 	Long:  "Testing my bois who all want to say up their homies and their neighborhood. Gang gang.",
// }

var rootCmd = &cobra.Command{
	Use:   "create-fs",
	Short: "This is going to, for now, just do an echo command",
	Long:  "This is going to for now just do an echo command",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("ayo!")
	// },
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.Execute()
	// gets appended to rootCmd

}

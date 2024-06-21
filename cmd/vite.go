/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// viteCmd represents the vite command
var viteCmd = &cobra.Command{
	Use:   "vite",
	Short: "Creates a vite app",
	Long:  `Accepts a name as a parameter to create a vite app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vite called")
		fmt.Println(args)

		install := exec.Command("bun", "create", "vite")
		output, err := install.Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Ran install!")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(viteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

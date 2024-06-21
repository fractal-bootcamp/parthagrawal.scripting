/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// viteCmd represents the vite command
var viteCmd = &cobra.Command{
	Use:   "vite",
	Short: "Creates a React + Typescript vite app",
	Long:  `Accepts parameters name as a parameter to create a vite app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vite called")
		fmt.Println(args)
		if len(args) == 0 {
			fmt.Println("You gotta specify a title bro")
			os.Exit(0)

		}

		title := args[0]

		install := exec.Command("bun", "create", "vite", title, "--template", "react-ts")
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

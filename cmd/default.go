/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sickodev/grandmaster-llm/engine"
	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Load simulation with default settings",
	Long:  `Loads the simulation with the default settings - OpenAI vs Gemini`,
	Run: func(cmd *cobra.Command, args []string) {
		blackFlag, _ := cmd.Flags().GetString("black")
		whiteFlag, _ := cmd.Flags().GetString("white")

		if whiteFlag == "" && blackFlag == "" {
			engine.StartDefault()
		}
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	defaultCmd.Flags().String("black", "", "The model that would play as the black side of the board.")
	defaultCmd.Flags().String("white", "", "The model that would play as the white side of the board.")
}

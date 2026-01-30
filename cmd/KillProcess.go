/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// KillProcessCmd represents the KillProcess command
var KillProcessCmd = &cobra.Command{
	Use:   "KillProcess",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("KillProcess called")
	},
}

func init() {
	rootCmd.AddCommand(KillProcessCmd)
}

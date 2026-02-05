/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"zp/internals/config"
	"zp/internals/loadebpf"
	"zp/internals/process"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zp",
	Short: "port scanner cli",
	Long:  `zp is a cli tool to scan open ports on the local machine using eBPF technology.`,
	Run: func(cmd *cobra.Command, args []string) {
		spec, err := loadebpf.LoadEBPFProgram()
		if err != nil {
			println("Error loading eBPF program:", err.Error())
			return
		}
		db := config.NewDbConfig("./badgerdb")
		proc := process.NewIdentifyProcess(3000, spec, db)
		for {
			result := proc.Identify()
			if result != nil {
				println("Event data:", result.Pid, result.Port)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "This is a version command",
	Long: `This is a version command. It prints the version of the aws-decrypter.`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("aws-decrypter v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

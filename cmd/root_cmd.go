package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "indy-build",
	Short: "indy-build is a cli tool to start build against indy",
	Long:  "indy-build is a cli client to start build against indy",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute executes the indy command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(mvnCmd)
	rootCmd.AddCommand(npmCmd)
}

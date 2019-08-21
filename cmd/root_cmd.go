package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "indy-build",
	Short: "indy-client is a cli client to start build against indy",
	Long:  "indy-client is a cli client to start build against indy",
	Run: func(cmd *cobra.Command, args []string) {
		goals := []string{"clean", "compile"}
		ExecMvn(goals, "/home/gli/workspaces/java/projects/nos/partyline/pom.xml", "")
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

}

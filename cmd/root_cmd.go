package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/gli/indy-build/process"
)

var indyURL, prjPom, prjTag, buildName string

var rootCmd = &cobra.Command{
	Use:   "indy-build",
	Short: "indy-build is a cli tool to start build against indy",
	Long:  "indy-build is a cli client to start build against indy",
	Run: func(cmd *cobra.Command, args []string) {
		process.RunBuild(indyURL, prjPom, prjTag, buildName)
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
	rootCmd.Flags().StringVarP(&indyURL, "indy_url", "i", "", "indy url.")
	rootCmd.Flags().StringVarP(&prjPom, "prjPom", "p", "", "project pom.")
	rootCmd.Flags().StringVarP(&prjTag, "prjTag", "t", "", "project tag.")
	rootCmd.Flags().StringVarP(&buildName, "buildName", "b", "", "build name.")
}

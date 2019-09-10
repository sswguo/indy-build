package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/gli/indy-build/process"
)

var indyURL, gitURL, tag, branch, buildName string

var rootCmd = &cobra.Command{
	Use:   "indy-build",
	Short: "indy-build is a cli tool to start build against indy",
	Long:  "indy-build is a cli client to start build against indy",
	Run: func(cmd *cobra.Command, args []string) {
		readyToRun := true
		checkout, checkoutType, validC := getCheckout()
		validV := validateArgs()
		readyToRun = validC && validV
		if readyToRun {
			process.RunBuild(indyURL, gitURL, checkoutType, checkout, buildName)
		}
	},
}

// Execute executes the indy command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateArgs() (valid bool) {
	valid = true
	if strings.TrimSpace(indyURL) == "" {
		fmt.Print("indyURL can not be empty.\n\n")
		valid = false
	}
	if strings.TrimSpace(buildName) == "" {
		fmt.Print("buildName can not be empty.\n\n")
	}
	return valid
}

func getCheckout() (checkout, checkoutType string, valid bool) {
	valid = true
	if strings.TrimSpace(tag) != "" {
		checkout = tag
		checkoutType = "tag"
	} else if strings.TrimSpace(branch) != "" {
		checkout = branch
		checkoutType = "branch"
	} else {
		fmt.Print("Error: tag or branch must be specified at least one\n\n")
		valid = false
	}
	return checkout, checkoutType, valid
}

func init() {
	rootCmd.Flags().StringVarP(&indyURL, "indy_url", "i", "", "indy url.")
	rootCmd.Flags().StringVarP(&gitURL, "gitURL", "g", "", "project git.")
	rootCmd.Flags().StringVarP(&tag, "tag", "t", "", "project git tag to build")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "project git branch to build.")
	rootCmd.Flags().StringVarP(&buildName, "buildName", "n", "", "build name.")
}

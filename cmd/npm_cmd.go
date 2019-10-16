package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/gli/indy-build/process"
)

var indyURL_npm, gitURL_npm, tag_npm, branch_npm, buildName_npm string

var npmCmd = &cobra.Command{
	Use:   "npm",
	Short: "do npm build against indy",
	Long:  "npm build against indy, includes build, folo, promote",
	Run: func(cmd *cobra.Command, args []string) {
		npmArgs := &baseArgs{
			indyURL:   indyURL_npm,
			gitURL:    gitURL_npm,
			tag:       tag_npm,
			branch:    branch_npm,
			buildName: buildName_npm,
		}
		readyToRun := true
		checkout, checkoutType, validC := getCheckout(npmArgs)
		validV := validateBaseArgs(npmArgs)
		validPrepare := process.CheckPrerequisites(CMD_NPM)
		readyToRun = validC && validV && validPrepare
		indyURL, gitURL, buildName := npmArgs.indyURL, npmArgs.gitURL, npmArgs.buildName
		if readyToRun {
			process.RunBuild(indyURL, gitURL, checkoutType, checkout, process.TYPE_NPM, buildName)
		}
	},
}

func init() {
	npmCmd.Flags().StringVarP(&indyURL_npm, "indy_url", "i", "", "indy url.")
	npmCmd.Flags().StringVarP(&gitURL_npm, "gitURL", "g", "", "project git.")
	npmCmd.Flags().StringVarP(&tag_npm, "tag", "t", "", "project git tag to build")
	npmCmd.Flags().StringVarP(&branch_npm, "branch", "b", "", "project git branch to build.")
	npmCmd.Flags().StringVarP(&buildName_npm, "buildName", "n", "", "build name.")

	npmCmd.MarkFlagRequired("indy_url")
	npmCmd.MarkFlagRequired("buildName")
	npmCmd.MarkFlagRequired("gitURL")
}

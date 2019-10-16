package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/gli/indy-build/process"
)

var indyURL_mvn, gitURL_mvn, tag_mvn, branch_mvn, buildName_mvn string

var mvnCmd = &cobra.Command{
	Use:   "maven",
	Short: "do maven build against indy",
	Long:  "maven build against indy, includes build, folo, promote",
	Run: func(cmd *cobra.Command, vArgs []string) {
		mvnArgs := &baseArgs{
			indyURL:   indyURL_mvn,
			gitURL:    gitURL_mvn,
			tag:       tag_mvn,
			branch:    branch_mvn,
			buildName: buildName_mvn,
		}
		readyToRun := true
		checkout, checkoutType, validC := getCheckout(mvnArgs)
		validV := validateBaseArgs(mvnArgs)
		validPrepare := process.CheckPrerequisites(CMD_MVN)
		readyToRun = validC && validV && validPrepare
		indyURL, gitURL, buildName := mvnArgs.indyURL, mvnArgs.gitURL, mvnArgs.buildName
		if readyToRun {
			process.RunBuild(indyURL, gitURL, checkoutType, checkout, process.TYPE_MVN, buildName)
		}
	},
}

func init() {
	mvnCmd.Flags().StringVarP(&indyURL_mvn, "indy_url", "i", "", "indy url.")
	mvnCmd.Flags().StringVarP(&gitURL_mvn, "gitURL", "g", "", "project git.")
	mvnCmd.Flags().StringVarP(&tag_mvn, "tag", "t", "", "project git tag to build")
	mvnCmd.Flags().StringVarP(&branch_mvn, "branch", "b", "", "project git branch to build.")
	mvnCmd.Flags().StringVarP(&buildName_mvn, "buildName", "n", "", "build name.")

	mvnCmd.MarkFlagRequired("indy_url")
	mvnCmd.MarkFlagRequired("buildName")
	mvnCmd.MarkFlagRequired("gitURL")
}

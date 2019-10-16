package cmd

import (
	"fmt"
	"strings"
)

const (
	CMD_MVN = "mvn"
	CMD_NPM = "npm"
)

type baseArgs struct {
	indyURL   string
	gitURL    string
	tag       string
	branch    string
	buildName string
}

func getCheckout(args *baseArgs) (checkout, checkoutType string, valid bool) {
	valid = true
	if strings.TrimSpace(args.tag) != "" {
		checkout = args.tag
		checkoutType = "tag"
	} else if strings.TrimSpace(args.branch) != "" {
		checkout = args.branch
		checkoutType = "branch"
	} else {
		fmt.Print("Error: tag or branch must be specified at least one\n\n")
		valid = false
	}
	return checkout, checkoutType, valid
}

func validateBaseArgs(bArgs *baseArgs) (valid bool) {
	valid = true
	if strings.TrimSpace(bArgs.indyURL) == "" {
		fmt.Print("indyURL can not be empty.\n\n")
		valid = false
	}
	if strings.TrimSpace(bArgs.buildName) == "" {
		fmt.Print("buildName can not be empty.\n\n")
	}
	return valid
}

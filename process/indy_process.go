package process

import (
	"fmt"
	"strings"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func prepareRepos(indyURL string, buildName string) {
	prepareHosted(indyURL, buildName)
	prepareGroup(indyURL, buildName)
}

func prepareHosted(indyURL string, buildName string) {
	hostedVars := template.IndyHostedVars{
		Name: buildName,
	}

	URL := fmt.Sprintf("%s/api/admin/stores/maven/hosted/%s", indyURL, buildName)

	hosted := template.IndyHostedTemplate(&hostedVars)
	fmt.Printf("Start creating hosted repo %s\n", buildName)
	result := putRequest(URL, strings.NewReader(hosted))
	if result {
		fmt.Printf("Hosted repo %s created successfully, check %s for details\n", buildName, URL)
	}
}

func prepareGroup(indyURL string, buildName string) {
	groupVars := template.IndyGroupVars{
		Name:         buildName,
		Constituents: []string{fmt.Sprintf("maven:hosted:%s", buildName), "maven:hosted:pnc-builds", "maven:remote:central"},
	}
	group := template.IndyGroupTemplate(&groupVars)

	URL := fmt.Sprintf("%s/api/admin/stores/maven/group/%s", indyURL, buildName)

	fmt.Printf("Start creating group repo %s\n", buildName)
	result := putRequest(URL, strings.NewReader(group))
	if result {
		fmt.Printf("Group repo %s created successfully, check %s for details\n", buildName, URL)
	}
}

func destroyRepos(indyURL string, buildName string) {
	destroyGroup(indyURL, buildName)
	// destroyHosted(indyURL, buildName)
}

func destroyHosted(indyURL string, buildName string) {
	URL := fmt.Sprintf("%s/api/admin/stores/maven/hosted/%s", indyURL, buildName)
	fmt.Printf("Start deleting hosted repo %s\n", buildName)
	result := delRequest(URL)
	if result {
		fmt.Printf("Hosted repo %s deleted successfully\n", buildName)
	}
}

func destroyGroup(indyURL string, buildName string) {
	URL := fmt.Sprintf("%s/api/admin/stores/maven/group/%s", indyURL, buildName)
	fmt.Printf("Start deleting group repo %s\n", buildName)
	result := delRequest(URL)
	if result {
		fmt.Printf("Group repo %s deleted successfully\n", buildName)
	}
}

func promote(indyURL, source, target string) {
	promoteVars := template.IndyPromoteVars{
		Source: source,
		Target: target,
	}
	promote := template.IndyPromoteJSONTemplate(&promoteVars)

	URL := fmt.Sprintf("%s/api/promotion/paths/promote", indyURL)

	fmt.Printf("Start promote request:\n %s\n\n", promote)
	respText, result := postRequest(URL, strings.NewReader(promote))

	if result {
		fmt.Printf("Promote successfully. Result is:\n %s\n\n", respText)
	} else {
		fmt.Printf("Promote failed. Result is:\n %s\n\n", respText)
	}
}

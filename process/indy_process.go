package process

import (
	"encoding/json"
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

func sealFolo(indyURL, foloId string) bool {
	URL := fmt.Sprintf("%s/api/folo/admin/%s/record", indyURL, foloId)
	fmt.Printf("Start to seal folo tracking: %s", foloId)
	_, result := postRequest(URL, nil)
	if result {
		fmt.Printf("Folo tracking sealing done: %s", foloId)
	} else {
		fmt.Printf("Folo tracking sealing failed: %s", foloId)
		return false
	}
	return true
}

func getFolo(indyURL, foloId string) ([]string, bool) {
	URL := fmt.Sprintf("%s/api/folo/admin/%s/record", indyURL, foloId)
	fmt.Printf("Start to get folo tracking: %s", foloId)
	data, result := getRequest(URL)
	if !result {
		fmt.Printf("Get folo tracking failed: %s", foloId)
		return nil, false
	}
	trackingContent := &TrackingContent{}
	err := json.Unmarshal([]byte(data), trackingContent)
	if err != nil {
		fmt.Printf("Get folo tracking failed: %s, Reason: %s ", foloId, err)
		return nil, false
	}
	upds := trackingContent.Uploads
	paths := make([]string, len(upds))
	for _, upd := range trackingContent.Uploads {
		paths = append(paths, upd.Path)
	}
	return paths, true
}

func promote(indyURL, source, target string, paths []string) {
	promoteVars := template.IndyPromoteVars{
		Source: source,
		Target: target,
		Paths:  paths,
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

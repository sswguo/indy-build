package process

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func prepareIndyRepos(indyURL, buildName string, buildMeta BuildMetadata) bool {
	if preapareIndyTargetHosted(indyURL, buildMeta) {
		if prepareIndyHosted(indyURL, buildMeta.buildType, buildName) {
			if prepareIndyGroup(indyURL, buildName, buildMeta) {
				return true
			}
		}
	}
	return false
}

func preapareIndyTargetHosted(indyURL string, buildMeta BuildMetadata) bool {
	buildType, target := buildMeta.buildType, buildMeta.promoteTarget

	URL := fmt.Sprintf("%s/api/admin/stores/%s/hosted/%s", indyURL, buildType, target)
	_, result := getRequest(URL)
	if result {
		fmt.Printf("Target hosted %s already exists, will bypass creation\n\n", target)
		return result
	}
	fmt.Printf("Target hosted %s not exists, will create first\n\n", target)

	hostedVars := template.IndyHostedVars{
		Name: target,
	}

	hosted := template.IndyHostedTemplate(&hostedVars)
	fmt.Printf("Start creating hosted repo %s\n", target)
	result = putRequest(URL, strings.NewReader(hosted))
	if result {
		fmt.Printf("Hosted repo %s created successfully, check %s for details\n", target, URL)
	} else {
		fmt.Printf("Hosted repo %s created failed, no following operations\n", target)
	}

	return result
}

func prepareIndyHosted(indyURL, buildType, buildName string) bool {
	hostedVars := template.IndyHostedVars{
		Name: buildName,
		Type: buildType,
	}

	URL := fmt.Sprintf("%s/api/admin/stores/%s/hosted/%s", indyURL, buildType, buildName)

	hosted := template.IndyHostedTemplate(&hostedVars)
	fmt.Printf("Start creating hosted repo %s\n", buildName)
	result := putRequest(URL, strings.NewReader(hosted))
	if result {
		fmt.Printf("Hosted repo %s created successfully, check %s for details\n", buildName, URL)
	} else {
		fmt.Printf("Hosted repo %s creation fail, no following operations\n", buildName)
	}
	return result
}

func prepareIndyGroup(indyURL, buildName string, buildMeta BuildMetadata) bool {
	buildType, target, central := buildMeta.buildType, buildMeta.promoteTarget, buildMeta.centralName
	groupVars := template.IndyGroupVars{
		Name:         buildName,
		Type:         buildMeta.buildType,
		Constituents: []string{fmt.Sprintf("%s:hosted:%s", buildType, buildName), fmt.Sprintf("%s:hosted:%s", buildType, target), fmt.Sprintf("%s:remote:%s", buildType, central)},
	}
	group := template.IndyGroupTemplate(&groupVars)

	URL := fmt.Sprintf("%s/api/admin/stores/%s/group/%s", indyURL, buildType, buildName)

	fmt.Printf("Start creating group repo %s\n", buildName)
	result := putRequest(URL, strings.NewReader(group))
	if result {
		fmt.Printf("Group repo %s created successfully, check %s for details\n", buildName, URL)
	} else {
		fmt.Printf("Group repo %s created failed, no following operations\n", buildName)
	}
	return result
}

func destroyIndyRepos(indyURL, buildType, buildName string) {
	destroyIndyGroup(indyURL, buildType, buildName)
	// destroyHosted(indyURL, buildName)
}

func destroyIndyHosted(indyURL, buildType, buildName string) {
	URL := fmt.Sprintf("%s/api/admin/stores/%s/hosted/%s", indyURL, buildType, buildName)
	fmt.Printf("Start deleting hosted repo %s\n", buildName)
	result := delRequest(URL)
	if result {
		fmt.Printf("Hosted repo %s deleted successfully\n", buildName)
	}
}

func destroyIndyGroup(indyURL, buildType, buildName string) {
	URL := fmt.Sprintf("%s/api/admin/stores/%s/group/%s", indyURL, buildType, buildName)
	fmt.Printf("Start deleting group repo %s\n", buildName)
	result := delRequest(URL)
	if result {
		fmt.Printf("Group repo %s deleted successfully\n", buildName)
	}
}

func sealIndyFolo(indyURL, foloId string) bool {
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

func getIndyFolo(indyURL, foloId string) ([]string, bool) {
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
		fmt.Printf("Promote Done. Result is:\n %s\n\n", respText)
	} else {
		fmt.Printf("Promote Error. Result is:\n %s\n\n", respText)
	}
}

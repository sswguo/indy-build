package template

import (
	"fmt"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func ExampleMvnSettingsTemplate() {
	settingsVar := template.MvnSettingsVars{
		LocalRepo: "/home/user/.m2/repo",
		BuildGroup: "build-1",
		IndyURL:    "http://indy.yourdomain.com",
	}
	settings := template.MvnSettingsTemplate(&settingsVar)
	fmt.Println(settings)
}

func ExampleIndyGroupTemplate() {
	groupVars := template.IndyGroupVars{
		Name:         "build-1",
		Constituents: []string{"maven:remote:central", "maven:hosted:build-1"},
	}
	group := template.IndyGroupTemplate(&groupVars)
	fmt.Println(group)
}

func ExampleIndyHostedTemplate() {
	hostedVars := template.IndyHostedVars{
		Name: "build-1",
	}

	hosted := template.IndyHostedTemplate(&hostedVars)
	fmt.Println(hosted)
}

func ExampleIndyPromoteJSONTemplate() {
	promoteVars := template.IndyPromoteVars{
		Source: "maven:hosted:build-1",
		Target: "maven:hosted:pnc-builds",
		Paths:  []string{"/a/b/c", "/x/y/z"},
	}
	promote := template.IndyPromoteJSONTemplate(&promoteVars)
	fmt.Println(promote)
}

package main

import (
	"fmt"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func main() {
	// cmd.Execute()
	groupVars := template.IndyGroupVars{
		Name:         "build-1",
		Constituents: []string{"maven:remote:central", "maven:hosted:build-1"},
	}

	s := template.IndyGroupTemplate(&groupVars)

	fmt.Println(s)
}

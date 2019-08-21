package main

import (
	"fmt"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func main() {
	// cmd.Execute()
	hostedVars := template.IndyHostedVars{
		Name: "build-1",
	}

	s := template.IndyHostedTemplate(&hostedVars)

	fmt.Println(s)
}

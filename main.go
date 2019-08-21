package main

import (
	"fmt"

	"gitlab.cee.redhat.com/gli/indy-build/cmd"
)

func main() {
	// cmd.Execute()
	settingsVar := cmd.MvnSettingsVars{
		UserHome:   "/home/gli",
		BuildGroup: "build-1",
		IndyUrl:    "http://indy-stable-next-devel.psi.redhat.com",
	}

	settings := cmd.MvnSettingsTemplate(&settingsVar)
	fmt.Println(settings)
}

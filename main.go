package main

import (
	"gitlab.cee.redhat.com/gli/indy-build/cmd"
)

func main() {
	cmd.Execute()
	// promoteVars := template.IndyPromoteVars{
	// 	Source: "maven:hosted:build-1",
	// 	Target: "maven:hosted:pnc-builds",
	// 	Paths:  []string{"/a/b/c", "/x/y/z"},
	// }
	// promote := template.IndyPromoteJSONTemplate(&promoteVars)
	// fmt.Println(promote)
}

package process

import (
	"fmt"
	"os/exec"
)

func runNpmBuild(indyURL, buildName string) {
	localRepo := fmt.Sprintf("/tmp/repo-%s", buildName)
	runNpmInstall(indyURL, localRepo, buildName)
	destroyRepo(localRepo)
}

func runNpmInstall(indyURL, prjLoc, buildName string) {
	runNpmCmd("install", indyURL, prjLoc, buildName)
}

func runNpmPublish(indyURL, prjLoc, buildName string) {
	runNpmCmd("publish", indyURL, prjLoc, buildName)
}

func runNpmCmd(cmd, indyURL, prjLoc, buildName string) {
	args := make([]string, 0)
	args = append(args, cmd)
	args = append(args, "--registry")
	registry := fmt.Sprintf("%s/api/folo/track/%s/npm/group/%s", indyURL, buildName, buildName)
	args = append(args, registry)
	finalCmd := "npm"
	fmt.Printf("Start executing: %s, in %s\n\n", getWholeCmd(finalCmd, args), prjLoc)
	exeCmd := exec.Command(finalCmd, args...)
	exeCmd.Dir = prjLoc
	printRealCmdOutput(exeCmd)
}

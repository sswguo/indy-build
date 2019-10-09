package process

import (
	"fmt"
	"os/exec"
	"path"
)

func checkCmd(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("Error: Can not find build command \"%s\", it is needed to run this build\n", cmd)
		return false
	}
	return true
}

func CheckPrerequisites(cmd string) bool {
	return checkCmd(cmd)
}

func RunBuild(indyURL, gitURL, checkoutType, checkout, buildName string) {
	dir := GetSrc(gitURL, checkout, checkoutType)
	prjPom := path.Join(dir, "pom.xml")
	if prepareRepos(indyURL, buildName) {
		runMvnBuild(indyURL, prjPom, buildName)
		sealed := sealFolo(indyURL, buildName)
		if sealed {
			paths, done := getFolo(indyURL, buildName)
			if done {
				promote(indyURL, fmt.Sprintf("maven:hosted:%s", buildName), "maven:hosted:pnc-builds", paths)
			}
		}

		destroyRepos(indyURL, buildName)
	}

	rmRepo(dir)
}

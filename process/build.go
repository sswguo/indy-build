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

func RunBuild(indyURL, gitURL, checkoutType, checkout, buildType, buildName string) {
	dir := GetSrc(gitURL, checkout, checkoutType)
	prjPom := path.Join(dir, "pom.xml")
	buildMeta := decideMeta(buildType)
	if buildMeta != nil {
		if prepareIndyRepos(indyURL, buildName, *buildMeta) {
			if buildType == TYPE_MVN {
				runMvnBuild(indyURL, prjPom, buildName)
			} else if buildType == TYPE_NPM {
				runNpmBuild(indyURL, buildName)
			}
			sealed := sealIndyFolo(indyURL, buildName)
			if sealed {
				paths, done := getIndyFolo(indyURL, buildName)
				if done {
					promote(indyURL, fmt.Sprintf("%s:hosted:%s", buildType, buildName), fmt.Sprintf("%s:hosted:%s", buildType, buildMeta.promoteTarget), paths)
				}
			}
			destroyIndyRepos(indyURL, buildType, buildName)
		}

	}

	rmRepo(dir)
}

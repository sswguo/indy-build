package process

import "fmt"

func RunBuild(indyURL, prjPom, prjTag, buildName string) {
	prepareRepos(indyURL, buildName)
	runMvnBuild(indyURL, prjPom, buildName)
	promote(indyURL, fmt.Sprintf("maven:hosted:%s", buildName), "maven:hosted:pnc-builds")
	destroyRepos(indyURL, buildName)
}

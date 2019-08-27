package process

import "fmt"

func RunBuild(indyURL, prjPom, prjTag, buildName string) {
	prepareRepos(indyURL, buildName)
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

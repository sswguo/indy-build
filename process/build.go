package process

import (
	"fmt"
	"path"
)

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

package process

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gitlab.cee.redhat.com/gli/indy-build/template"
)

func runMvnBuild(indyURL, prjPom, buildName string) {
	localRepo := fmt.Sprintf("/tmp/repo-%s", buildName)
	settings := prepareMvnSettings(indyURL, localRepo, buildName)
	runMvn([]string{"clean", "deploy"}, prjPom, settings)
	destroyMvnSettings(settings)
	destroyRepo(localRepo)
}

func runMvn(goals []string, pomfile string, settingsFile string) {
	args := make([]string, 0)
	for _, goal := range goals {
		args = append(args, goal)
	}

	args = append(args, "-DskipTests")

	if len(strings.Trim(pomfile, "")) > 0 {
		args = append(args, "-f")
		args = append(args, pomfile)
	}
	if len(strings.Trim(settingsFile, "")) > 0 {
		args = append(args, "-s")
		args = append(args, settingsFile)
	}
	fmt.Printf("Start executing: %s\n\n", getWholeCmd("mvn", args))
	printRealCmdOutput(exec.Command("mvn", args...))
}

func prepareMvnSettings(IndyURL, localRepo, buildName string) string {
	var repo string
	if strings.TrimSpace(localRepo) != "" {
		repo = strings.TrimSpace(localRepo)
	} else {
		userHome := os.Getenv("HOME")
		if strings.TrimSpace(userHome) == "" {
			userHome = "~"
		}
		repo = fmt.Sprintf("%s/.m2/%s", userHome, buildName)
	}
	settingsVar := template.MvnSettingsVars{
		LocalRepo:  repo,
		BuildGroup: buildName,
		IndyURL:    IndyURL,
	}
	settings := template.MvnSettingsTemplate(&settingsVar)

	tmp := os.Getenv("TMPDIR")
	if strings.TrimSpace(tmp) == "" {
		tmp = "/tmp"
	}
	settingsFile := fmt.Sprintf("%s/settings-%s.xml", tmp, buildName)

	storeFile(settingsFile, settings)

	fmt.Printf("settings generated: %s\n\n", settingsFile)
	fmt.Printf("settings content: %s\n\n", settings)

	return settingsFile
}

func destroyMvnSettings(settingsFile string) {
	os.Remove(settingsFile)
	fmt.Printf("settings removed: %s\n", settingsFile)
}

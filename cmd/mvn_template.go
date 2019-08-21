package cmd

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type MvnSettingsVars struct {
	UserHome, BuildGroup, IndyUrl string
}

// Usage Example:
//
// settingsVar := cmd.MvnSettingsVars{
// 		UserHome:   "/home/user",
// 		BuildGroup: "build-1",
// 		IndyUrl:    "http://indy.yourdomain.com",
// 	}
// settings := cmd.MvnSettingsTemplate(&settingsVar)
func MvnSettingsTemplate(settingsVars *MvnSettingsVars) string {
	var settingsTemplate = `
<settings>
  <localRepository>{{.UserHome}}/.m2/repo-{{.BuildGroup}}</localRepository>
  <mirrors>
    <mirror>
      <id>indy-build</id>
      <mirrorOf>*</mirrorOf>
      <url>{{.IndyUrl}}/api/content/maven/group/{{.BuildGroup}}</url>
    </mirror>
  </mirrors>
  <profiles>
    <profile>
      <id>resolve-settings</id>
      <repositories>
        <repository>
          <id>central</id>
          <url>{{.IndyUrl}}/api/content/maven/group/indy-build</url>
          <releases>
            <enabled>true</enabled>
          </releases>
          <snapshots>
            <enabled>true</enabled>
          </snapshots>
        </repository>
      </repositories>
      <pluginRepositories>
        <pluginRepository>
          <id>central</id>
          <url>{{.IndyUrl}}/api/content/maven/group/indy-build</url>
          <releases>
            <enabled>true</enabled>
          </releases>
          <snapshots>
            <enabled>true</enabled>
          </snapshots>
        </pluginRepository>
      </pluginRepositories>
    </profile>
    
    <profile>
      <id>deploy-settings</id>
      <properties>
        <altDeploymentRepository>{{.BuildGroup}}::default::{{.IndyUrl}}/api/content/maven/group/{{.BuildGroup}}</altDeploymentRepository>
      </properties>
    </profile>
    
  </profiles>
  <activeProfiles>
    <activeProfile>resolve-settings</activeProfile>
    <activeProfile>deploy-settings</activeProfile>
  </activeProfiles>
</settings>
`
	t := template.Must(template.New("settings").Parse(settingsTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, settingsVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

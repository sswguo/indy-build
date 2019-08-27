package template

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

// MvnSettingsVars ...
type MvnSettingsVars struct {
	LocalRepo, BuildGroup, IndyURL string
}

// MvnSettingsTemplate ...
func MvnSettingsTemplate(settingsVars *MvnSettingsVars) string {
	var settingsTemplate = `<settings>
  <localRepository>{{.LocalRepo}}</localRepository>
  <mirrors>
    <mirror>
      <id>indy-build</id>
      <mirrorOf>*</mirrorOf>
      <url>{{.IndyURL}}/api/folo/track/{{.BuildGroup}}/maven/group/{{.BuildGroup}}</url>
    </mirror>
  </mirrors>
  <profiles>
    <profile>
      <id>resolve-settings</id>
      <repositories>
        <repository>
          <id>central</id>
          <url>{{.IndyURL}}/api/folo/track/{{.BuildGroup}}/maven/group/{{.BuildGroup}}</url>
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
          <url>{{.IndyURL}}/api/folo/track/{{.BuildGroup}}/maven/group/{{.BuildGroup}}</url>
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
        <altDeploymentRepository>{{.BuildGroup}}::default::{{.IndyURL}}/api/folo/track/{{.BuildGroup}}/maven/group/{{.BuildGroup}}</altDeploymentRepository>
      </properties>
    </profile>
    
  </profiles>
  <activeProfiles>
    <activeProfile>resolve-settings</activeProfile>
    <activeProfile>deploy-settings</activeProfile>
  </activeProfiles>
</settings>`

	t := template.Must(template.New("settings").Parse(settingsTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, settingsVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

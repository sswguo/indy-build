package template

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type IndyGroupVars struct {
	Name         string
	Constituents []string
}

func IndyGroupTemplate(indyGroupVars *IndyGroupVars) string {
	groupTemplate := `
	{
		"type" : "group",
		"key" : "maven:group:{{.Name}}",
		"metadata" : {
			"changelog" : "init group {{.Name}}"
		},
		"disabled" : false,
		"constituents" : [{{range $con := .Constituents}}"{{$element}}",{{end}}],
		"packageType" : "maven",
		"name" : "{{.Name}}",
		"type" : "group",
		"disable_timeout" : 0,
		"path_style" : "plain",
		"authoritative_index" : false,
		"prepend_constituent" : false
	}
	`
	t := template.Must(template.New("settings").Parse(groupTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, indyGroupVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

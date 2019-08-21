package template

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

// IndyPromoteVars ...
type IndyPromoteVars struct {
	Source string
	Target string
	Paths  []string
}

// IndyPromoteJSONTemplate ...
func IndyPromoteJSONTemplate(indyPromoteVars *IndyPromoteVars) string {
	request := `{
  "async": false,
  "source": "{{.Source}}",
  "target": "{{.Target}}",
  {{if gt (len .Paths) 0}}
  "paths": [{{range $index,$path := .Paths}}"{{$path}}"{{if isNotLast $index $.Paths}},{{end}}{{end}}],
  {{end}}
  "purgeSource": false,
  "dryRun": false,
  "fireEvents": true,
  "failWhenExists": true
}`

	t := template.Must(template.New("settings").Funcs(isNotLast).Parse(request))
	var buf bytes.Buffer
	err := t.Execute(&buf, indyPromoteVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

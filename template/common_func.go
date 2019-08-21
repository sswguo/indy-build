package template

import "text/template"

var isNotLast = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"isNotLast": func(index int, array []string) bool {
		return index < len(array)-1
	},
}

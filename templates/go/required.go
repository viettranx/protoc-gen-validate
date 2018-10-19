package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			{{ err . "value is required" }}
		}
	{{ end }}
`

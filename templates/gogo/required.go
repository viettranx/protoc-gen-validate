package gogo

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		{{ if .Gogo.Nullable }}
			if {{ accessor . }} == nil {
				{{ err . "value is required" }}
			}
		{{ end }}
	{{ end }}
`

package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					{{ err . "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					{{ err . "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					{{ err . "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					{{ err . "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				{{ err . "value must be less than " $r.GetLt }}
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					{{ err . "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					{{ err . "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					{{ err . "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					{{ err . "value must be outside range (" $r.GetLte ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				{{ err . "value must be less than or equal to " $r.GetLte }}
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			{{ err . "value must be greater than " $r.GetGt }}
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			{{ err . "value must be greater than or equal to " $r.GetGte }}
		}
	{{ end }}
`

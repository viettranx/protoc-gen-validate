package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if len({{ accessor . }}) != {{ $r.GetLen }} {
				{{ err . "value length must be " $r.GetLen " bytes" }}
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinLen }} {
				{{ err . "value length must be " $r.GetMinLen " bytes" }}
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				{{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " bytes, inclusive" }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinLen }} {
				{{ err . "value length must be at least " $r.GetMinLen " bytes" }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			{{ err . "value length must be at most " $r.GetMaxLen " bytes" }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			{{ err . "value does not have prefix " (byteStr $r.GetPrefix) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			{{ err . "value does not have suffix " (byteStr $r.GetSuffix) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			{{ err . "value does not contain " (byteStr $r.GetContains) }}
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; !ok {
			{{ err . "value must be in list " $r.In }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; ok {
			{{ err . "value must not be in list " $r.NotIn }}
		}
	{{ end }}

	{{ if $r.Const }}
		if !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			{{ err . "value must equal " $r.Const }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			{{ err . "value must be a valid IP address" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			{{ err . "value must be a valid IPv4 address" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			{{ err . "value must be a valid IPv6 address" }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		{{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
	}
	{{ end }}
`

{{define "back/status"}}
{{if eq .Status "A"}}
Active
{{else}}
Disabled
{{end}}
{{end}}
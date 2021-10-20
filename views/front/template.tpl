{{define "front/template"}}
<!DOCTYPE html>
<html lang="en">
<head>
	{{template "front/meta"}}
	{{template "front/links"}}
	<title>{{with .title}}{{.title}}{{else}}The Chan{{end}}</title>
</head>
<body>
	{{.page | asHTML}}

	{{template "front/scripts"}}
</body>
</html>
{{end}}
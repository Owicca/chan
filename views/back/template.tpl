{{define "back/template"}}
<!DOCTYPE html>
<html lang="en">
<head>
	{{template "back/meta" .}}
	{{template "back/links" .}}
	<title>{{with .title}}{{.title}}{{else}}The Chan{{end}}</title>
</head>
<body class="container-sm">
	<header>
		{{template "back/nav" .}}
	</header>
	<main>
		{{asHTML .page}}
	</main>
	<footer class="fixed-bottom">
	</footer>
	{{/*{{template "back/alerts" .}}*/}}
	{{template "back/scripts" .}}
</body>
</html>
{{end}}

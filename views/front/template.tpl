{{define "front/template"}}
<!DOCTYPE html>
<html lang="en">
<head>
	{{template "front/meta"}}
	{{template "front/links"}}
	<title>{{with .title}}{{.title}}{{else}}The Chan{{end}}</title>
</head>
<body>
	<header>
		{{template "front/nav" .}}
	</header>
	<main>
		{{.page | asHTML}}
	</main>
	<footer class="fixed-bottom">
	</footer>
	{{template "front/scripts"}}
</body>
</html>
{{end}}

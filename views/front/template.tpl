{{define "front/template"}}
<!DOCTYPE html>
<html lang="en">
<head>
	{{template "front/meta"}}
	{{template "front/links"}}
	<title>{{with .title}}{{.title}}{{else}}Imageboard{{end}}</title>
</head>
<body class="{{if .is_index}}is_index{{end}}{{if .is_thread}}is_thread{{end}}">
	<header>
		{{template "front/nav" .}}
	</header>
	<hr>
	<main>
		{{asHTML .page}}
	</main>
	<hr>
	<footer class="fixed-bottom">
		{{template "front/nav" .}}
	</footer>
	{{template "front/scripts"}}
</body>
</html>
{{end}}

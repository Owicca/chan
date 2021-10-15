{{define "back/template"}}
<!DOCTYPE html>
<html lang="en">
<head>
	{{template "back/meta"}}
	{{template "back/links"}}
	<title>{{.title}}</title>
</head>
<body class="container-sm">
	<header class="">
		{{template "back/nav"}}
	</header>
	<main class="mt-3">
		{{.page | asHTML}}
	</main>
	<footer class="fixed-bottom">
		footer
	</footer>
	{{template "back/scripts"}}
</body>
</html>
{{end}}
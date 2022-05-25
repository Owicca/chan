{{define "back/actions"}}
<ul>
	<li>
		{{if .view_name}}
			<a href="/admin/{{.view_name}}/{{.view_id}}/">View</a>
		{{end}}
		{{if .update_name}}
			<a href="/admin/{{.update_name}}/{{.update_id}}/">Update</a>
		{{end}}
	</li>
</ul>
{{end}}

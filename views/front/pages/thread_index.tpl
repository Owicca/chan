{{define "front/thread_index"}}
<ul>
{{range $idx, $post := .posts}}
	<li>{{template "post.tpl" $post}}</li>
{{end}}
</ul>
{{end}}
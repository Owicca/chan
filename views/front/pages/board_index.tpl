{{define "front/board_index"}}
<ul>
{{range $idx, $thread := .threads}}
	<li><a href="{{$thread.Path}}">{{$thread.Name}}</a></li>
{{else}}
	<li>No threads available!</li>
{{end}}
</ul>
{{end}}
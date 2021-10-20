{{define "back/index"}}
<ul>
	<li>{{.site.name}}</li>
	<li>{{.site.title}}</li>
	<li>{{.site.welcome}}</li>
</ul>
<ul>
{{range $col, $boards := .topics}}
	<li>
		{{$col}}
		<ul>
			{{range $board := $boards}}
				<p>{{$board.Name}}</p>
				<p>{{$board.Code}}</p>
				<p>{{$board.Description}}</p>
				<ul>
					{{range $media := $board.MediaList}}
					<img src="{{$media.Path}}">
					{{end}}
				</ul>
			{{end}}
		</ul>
	</li>
{{end}}
</ul>
{{end}}
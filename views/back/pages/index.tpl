{{define "back/index"}}
<ul class="container">
{{range $topic := .topic_list}}
	<li class="row">
		{{$topic.Name}}
		<ul class="col-2 container">
			{{range $board := $topic.BoardList}}
			<li class="row">
				<div class="card">
					{{range $media := $board.MediaList}}
					<img class="image rounded card-img-top" src="{{$media.Path}}">
					{{end}}
					<div class="card-body">
						<h5 class="card-title">{{$board.Name}}</h5>
						<p class="card-text">{{$board.Description}}</p>
						<a href="/admin/boards/{{$board.ID}}/" class="btn btn-primary">View</a>
					</div>
				</div>
			</li>
			{{end}}
		</ul>
	</li>
{{end}}
</ul>
{{end}}
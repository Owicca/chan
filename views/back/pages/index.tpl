{{define "back/index"}}
<ul class="container">
{{range $idx, $topic := .topic_list}}
	<li class="row row-topic al_left">
		<h3 class="">{{$topic.Name}}</h3>
		<ul class="container">
			<li class="row row--board">
				<table class="table table-sm table-striped align-middle">
					<thead>
						<th scope="col">#</th>
						<th scope="col">Status</th>
						<th scope="col">Image</th>
						<th scope="col">Name</th>
						<th scope="col">Description</th>
						<th scope="col">Actions</th>
					</thead>
					<tbody>
						{{range $board := $topic.BoardList}}
						<tr>
							<th scope="row">
								{{$board.ID}}
							</th>
							<td>
								{{if gt $board.Deleted_at 0}}
									Deleted at: {{unixToUTC $board.Deleted_at}}
								{{else}}
									Active
								{{end}}
							</td>
							<td>
								{{range $media := $board.MediaList}}
									<img class="image rounded card-img-top" src="{{$media.Path}}">
								{{end}}
							</td>
							<td>
								{{$board.Name}}
							</td>
							<td>
								{{$board.Description}}
							</td>
							<td>
								<a href="/admin/boards/{{$board.ID}}/" class="btn btn-primary">View</a>
							</td>
						</tr>
						{{else}}
						<tr>
							<td colspan="6">No boards available!</td>
						</tr>
						{{end}}
					</tbody>
				</table>
			</li>
		</ul>
	</li>
{{end}}
</ul>
{{end}}

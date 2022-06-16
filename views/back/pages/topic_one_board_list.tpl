{{define "back/topic_one_board_list"}}
<table class="table table-sm table-striped align-middle">
	<thead>
		<tr>
			<td scope="col">#</td>
			<td scope="col">Name</td>
			<td scope="col">Threads</td>
		</tr>
	</thead>
	<tbody>
{{range $board := .topic.BoardList}}
	<tr>
		<th scope="row">
			{{$board.ID}}
		</th>
		<td>
			<a href="/admin/boards/{{$board.ID}}/">{{$board.Name}}</a>
		</td>
		<td>
			{{$thread_count := len $board.ThreadList}}
			{{if gt $thread_count 0}}
				<a href="/admin/boards/{{$board.ID}}/threads/">
					{{$thread_count}}
				</a>
			{{else}}
				<span>{{$thread_count}}</span>
			{{end}}
		</td>
	</tr>
{{else}}
	<tr>
		<td colspan="4">No boards available!</td>
	</tr>
{{end}}
	</tbody>
</table>
{{end}}

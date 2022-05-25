{{define "back/board_list"}}
<a href="/admin/boards/add/">Add new</a>
<table class="table table-sm table-striped align-middle">
	<thead>
		<tr>
			<td scope="col">#</td>
			<td scope="col">Name</td>
			<td scope="col">Threads</td>
			<td scope="col">Actions</td>
		</tr>
	</thead>
	<tbody>
{{range $board := .board_list}}
	<tr>
		<th scope="row">
			{{$board.ID}}
		</th>
		<td>
			<a href="/admin/boards/{{$board.ID}}/">{{$board.Name}}</a>
		</td>
		<td>
			{{if gt $board.Thread_count 0}}
				<a href="/admin/boards/{{$board.ID}}/threads/">
					{{$board.Thread_count}}
				</a>
			{{else}}
				<span>{{$board.Thread_count}}</span>
			{{end}}
		</td>
		<td>
			{{template "back/actions" params "update_name" "boards" "update_id" $board.ID}}
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

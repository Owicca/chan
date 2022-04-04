{{define "back/board_list"}}
<a href="/admin/boards/add/">Add new</a>
<table>
	<thead>
		<tr>
			<td>ID</td>
			<td>Name</td>
			<td>Threads</td>
			<td></td>
		</tr>
	</thead>
	<tbody>
{{range $board := .board_list}}
	<tr>
		<td><a href="/admin/boards/{{$board.ID}}/">{{$board.ID}}</a></td>
		<td>{{$board.Name}}</td>
		<td>
			{{if gt $board.Thread_count 0}}
			<a href="/admin/boards/{{$board.ID}}/threads/">
				{{$board.Thread_count}}
			</a>
			{{else}}
			<span>{{$board.Thread_count}}</span>
			{{end}}
		</td>
		<td>{{template "back/actions" params "Name" "boards" "ID" $board.ID}}</td>
	</tr>
{{else}}
	<tr>
		<td>No boards available!</td>
	</tr>
{{end}}
	</tbody>
</table>
{{end}}
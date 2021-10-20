{{define "back/board_list"}}
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
{{range $board := .boards}}
	<tr>
		<td><a href="/admin/boards/{{$board.ID}}/">{{$board.ID}}</a></td>
		<td>{{$board.Name}}</td>
		<td>0</td>
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
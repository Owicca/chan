{{define "back/topic_list"}}
<a href="/admin/topics/add/">Add new</a>
<table class="table table-sm table-striped align-middle">
	<thead>
		<tr>
			<td scope="col">#</td>
			<td scope="col">Name</td>
			<td scope="col">Board count</td>
			<td scope="col">Deleted</td>
			<td scope="col">Actions</td>
		</tr>
	</thead>
	<tbody>
{{range $topic := .topic_list}}
	<tr>
		<th scope="row">
			{{$topic.ID}}
		</th>
		<td>
			<a href="/admin/topics/{{$topic.ID}}/">{{$topic.Name}}</a>
		</td>
		<td>
			{{$board_count := len $topic.BoardList}}
			{{if gt $board_count 0}}
				<a href="/admin/topics/{{$topic.ID}}/boards/">
					{{$board_count}}
				</a>
			{{else}}
				<span>{{$board_count}}</span>
			{{end}}
		</td>
		<td>
			{{if gt $topic.Deleted_at 0}}
				{{unixToUTC $topic.Deleted_at}}
			{{else}}
				Active
			{{end}}
		</td>
		<td>
			{{template "back/actions" params "update_name" "topics" "update_id" $topic.ID}}
		</td>
	</tr>
{{else}}
	<tr>
		<td>No topics available!</td>
	</tr>
{{end}}
	</tbody>
</table>
{{end}}

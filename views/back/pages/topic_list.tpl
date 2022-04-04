{{define "back/topic_list"}}
<a href="/admin/topics/add/">Add new</a>
<table>
	<thead>
		<tr>
			<td>ID</td>
			<td>Name</td>
			<td>Deleted</td>
			<td></td>
		</tr>
	</thead>
	<tbody>
{{range $topic := .topic_list}}
	<tr>
		<td><a href="/admin/topics/{{$topic.ID}}/">{{$topic.ID}}</a></td>
		<td>{{$topic.Name}}</td>
		<td>
			{{if gt $topic.Deleted_at 0}}
				{{unixToUTC $topic.Deleted_at}}
			{{else}}
				Active
			{{end}}
		</td>
		<td>{{template "back/actions" params "Name" "topics" "ID" $topic.ID}}</td>
	</tr>
{{else}}
	<tr>
		<td>No topics available!</td>
	</tr>
{{end}}
	</tbody>
</table>
{{end}}
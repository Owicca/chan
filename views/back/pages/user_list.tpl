{{define "back/user_list"}}
<table>
	<thead>
		<tr>
			<td>ID</td>
			<td>Name</td>
			<td>Role</td>
			<td>Email</td>
			<td>Status</td>
			<td></td>
		</tr>
	</thead>
	<tbody>
{{range $user := .users}}
	<tr>
		<td><a href="/admin/users/{{$user.ID}}/">{{$user.ID}}</a></td>
		<td>{{$user.Username}}</td>
		<td>{{$user.Role.Name}}</td>
		<td><a href="mailto:{{$user.Email}}">{{$user.Email}}</a></td>
		<td>{{template "back/status" params "Status" $user.Status}}</td>
		<td>{{template "back/actions" params "Name" "users" "ID" $user.ID}}</td>
	</tr>
{{else}}
	<tr>
		<td>No users available!</td>
	</tr>
{{end}}
	</tbody>
</table>
{{end}}
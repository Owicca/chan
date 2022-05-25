{{define "back/user_list"}}
<table class="table table-sm table-stripped align-middle text-start">
	<thead>
		<tr>
			<th scope="col">#</th>
			<th scope="col">Name</th>
			<th scope="col">Role</th>
			<th scope="col">Email</th>
			<th scope="col">Status</th>
			<th scope="col">Actions</th>
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
				<td>{{template "back/actions" params "update_name" "users" "update_id" $user.ID}}</td>
			</tr>
		{{else}}
			<tr>
				<td colspan="6">No users available!</td>
			</tr>
		{{end}}
	</tbody>
</table>
{{end}}

{{define "back/user"}}
{{if .user}}
<form method="POST" action="/admin/users/{{.user.ID}}/" class="">
	<input type="hidden" id="user_id" name="user_id" value="{{.user.ID}}" />
	<div class="input-group input-group-sm mb-3">
		<label for="username" class="input-group-text">Username: </label>
		<input type="text" id="username" class="form-control" name="username" value="{{.user.Username}}" />
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="email" class="input-group-text">Email: </label>
		<input type="text" id="email" class="form-control" name="email" value="{{.user.Email}}" />
	</div>
	<div class="input-group input-group-sm mb-3">
		<label class="input-group-text">Status: </label>
	{{range $status, $id := .statusList}}
		<div class="form-check d-flex justify-content-start me-3">
			<input type="radio" id="status_{{$id}}" class="form-check-input me-1" name="status" value="{{$id}}" {{if eq $id $.user.Status}}checked="checked"{{end}} />
			<label for="status_{{$id}}" class="form-check-label">{{$status}}</label>
		</div>
	{{else}}
		<p>no status</p>
	{{end}}
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="role" class="input-group-text">Role: </label>
		<select id="role" class="form-select" name="role">
			{{range $role := .roles}}
				<option
				value="{{$role.ID}}"
				{{if eq $role.ID $.user.Role.ID}}selected="selected"{{end}}
				>{{$role.Name}}</option>
			{{else}}
				<option disabled="true">No role available</option>
			{{end}}
		</select>
	</div>
	<div class="actions">
		<input type="submit" class="btn btn-success" value="Submit" />
		<a href="/admin/users/" class="btn btn-secondary">Back</a>
	</div>
</form>
{{else}}
	<p>User not found</p>
{{end}}
{{end}}
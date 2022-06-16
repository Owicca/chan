{{define "back/user_create_form"}}
<form class="text-left" method="POST" action="/admin/users/">
	<div class="input-group input-group-sm mb-3">
		<label for="email" class="input-group-text">Email</label>
		<input type="email" name="email" id="email" class="form-control" placeholder="email@email.com">
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="password1" class="input-group-text">Password</label>
		<input type="password" name="password1" id="password1" class="form-control" placeholder="password">
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="password2" class="input-group-text">Verify password</label>
		<input type="password" name="password2" id="password2" class="form-control" placeholder="password">
	</div>
	<div class="input-group input-group-sm mb-3">
			<label class="input-group-text">Role: </label>
			<select class="form-select" name="role" id="role">
			{{range $role := .user_role_list}}
					<option value="{{$role.ID}}">{{$role.Name}}</option>
			{{end}}
			</select>
	</div>
	<div class="input-group input-group-sm mb-3 d-none" id="boardCnt">
			<label class="input-group-text">Board: </label>
			<select class="form-select" name="board" id="board">
			{{range $board := .board_list}}
					<option value="{{$board.ID}}">{{$board.Name}}</option>
			{{end}}
			</select>
	</div>
	<div class="input-group input-group-sm mb-3">
			<label class="input-group-text">Status: </label>
			<select class="form-select" name="status">
			{{range $status, $val := .user_status_list}}
					<option value="{{$val}}">{{$status}}</option>
			{{end}}
			</select>
	</div>
	<input type="submit" class="btn btn-primary" value="Create">
</form>
{{end}}

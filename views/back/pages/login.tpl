{{define "back/login"}}
<div class="d-flex justify-content-center">
	<form class="text-left" method="POST" action="/admin/login/">
		<div class="form-floating mb-3">
			<label for="email" class="form-label">Email</label>
			<input type="email" name="email" id="email" class="form-control" placeholder="email@email.com">
		</div>
		<div class="form-floating mb-3">
			<label for="password1" class="form-label">Password</label>
			<input type="password" name="password1" id="password1" class="form-control" placeholder="password">
		</div>
		<div class="form-floating mb-3">
			<label for="password2" class="form-label">Verify password</label>
			<input type="password" name="password2" id="password2" class="form-control" placeholder="password">
		</div>
		<input type="submit" class="btn btn-primary" value="Log in">
	</form>
</div>
{{end}}

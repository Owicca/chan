{{define "back/nav"}}
<nav class="container-sm navbar navbar-light bg-light fixed-top">
	<a href="/admin/" class="navbar-brand">Chan</a>
	<button class="navbar-toggler" type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasNavbar" aria-controls="offcanvasNavbar">
		<span class="navbar-toggler-icon"></span>
	</button>
	<div class="offcanvas offcanvas-end" tabindex="-1" id="offcanvasNavbar">
		<div class="offcanvas-header">
			<button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close">Close</button>
		</div>
		<div class="offcanvas-body">
			<ul class="container navbar-nav">
				{{if (or (not .data.user) (and .data.user (eq .data.user.ID 0)))}}
					<li class="navbar-item"><a href="/admin/login/">Login</a></li>
				{{else}}
					<li class="navbar-item"><a href="/admin/">Home</a></li>
					<li class="navbar-item"><a href="/admin/users/">Users</a></li>
					<li class="navbar-item"><a href="/admin/topics/">Topics</a></li>
					<li class="navbar-item"><a href="/admin/boards/">Boards</a></li>
					<li class="navbar-item"><a href="/admin/threads/">Threads</a></li>
					<li class="navbar-item">
						<form method="POST" action="/admin/logout/">
							<input type="hidden" name="logout" value="1">
							<input type="submit" value="Logout">
						</form>
					</li>
				{{end}}
			</ul>
		</div>
	</div>
</nav>
{{end}}

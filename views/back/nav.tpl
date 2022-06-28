{{define "back/nav"}}
<nav class="container-sm navbar navbar-light bg-light fixed-top">
	<a href="/admin/" class="navbar-brand">Imageboard</a>
	<button class="navbar-toggler" type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasNavbar" aria-controls="offcanvasNavbar">
		<span class="navbar-toggler-icon"></span>
	</button>
	<div class="offcanvas offcanvas-end" tabindex="-1" id="offcanvasNavbar">
		<div class="offcanvas-header">
			<button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close">Close</button>
		</div>
		<div class="offcanvas-body">
			<ul class="container navbar-nav">
				{{/*{{if .data.user}}*/}}
					<li class="navbar-item"><a class="nav-link" href="/admin/">Home</a></li>
					<li class="navbar-item"><a class="nav-link" href="/admin/users/">Users</a></li>
					<li class="navbar-item"><a class="nav-link" href="/admin/topics/">Topics</a></li>
					<li class="navbar-item"><a class="nav-link" href="/admin/boards/">Boards</a></li>
					<li class="navbar-item"><a class="nav-link" href="/admin/threads/">Threads</a></li>
					<li class="navbar-item"><a class="nav-link" href="/admin/settings/">Settings</a></li>
					<li class="navbar-item">
							<form method="POST" action="/admin/logout/" class="nav-link">
								<input type="hidden" name="logout" value="1">
								<input type="submit" value="Logout">
							</form>
					</li>
				{{/*{{else}}
					<li class="navbar-item"><a href="/admin/login/">Login</a></li>
				{{end}}*/}}
			</ul>
		</div>
	</div>
</nav>
{{end}}

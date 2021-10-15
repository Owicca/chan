{{define "back/user_list"}}
<ul>
{{range $user := .users}}
	<li>
		<a href="/admin/users/{{$user.ID}}/">{{$user.Username}}</a>({{$user.Role.Name}}): <a href="mailto:{{$user.Email}}">{{$user.Email}}</a>
	</li>
{{else}}
	<li>No user available!</li>
{{end}}
</ul>
{{end}}
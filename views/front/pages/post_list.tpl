{{define "front/post_list"}}
{{template "front/create_reply_form" .}}
<hr>
{{template "front/post_list_nav_top" .}}
<hr>
<div class="board">
<div class="thread">

<ul>
{{range $idx, $post := .post_list}}
	{{$type := "reply"}}
	{{if eq $idx 0}}
		{{$type = "op"}}
	{{end}}

	{{$trp := ""}}
	{{if $post.SecureTripcode}}
		{{$trp = (printf "!!%s" $post.SecureTripcode)}}
	{{else if $post.Tripcode}}
		{{$trp = (printf "!%s" $post.Tripcode)}}
	{{end}}

	{{$pipe := (params "post" $post "type" $type "trp" $trp)}}
	{{template "front/post_one" $pipe}}
{{end}}
	<!-- <noscript> -->
{{template "front/create_reply_form_quick" .}}
	<!-- </noscript> -->
</ul>

</div>
</div>
<hr>
{{template "front/post_list_nav_bot" .}}
{{end}}

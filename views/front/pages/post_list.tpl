{{define "front/post_list"}}
{{$form_action := (printf "/boards/%s/threads/%s/" .board_code .thread_id)}}
{{$form_params := (params "form_action" $form_action "form_button_label" "Post a Reply")}}
{{template "front/create_reply_form" $form_params}}
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

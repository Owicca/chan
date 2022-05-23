{{define "front/thread_list"}}
{{template "front/create_reply_form" .}}
<hr>
{{template "front/post_list_nav_top" .}}
{{/* remove return, add search box instead */}}
{{/* remove post/media count */}}
<hr>
<div class="board">
{{range $t_idx, $thread := .thread_list}}
{{$previewLen := (len $thread.Preview)}}
{{if eq $previewLen 0}}
	{{continue}}
{{end}}
<div class="thread">

<ul>
{{range $idx, $post := $thread.Preview}}
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
	{{if eq $idx 0}}
		<li class="al_left">
			<div class="postContainer {{$type}}Container">
				<span><a href="/boards/{{$.board_code}}/threads/{{$thread.ID}}/">Click here</a> to view.</span>
			</div>
		</li>
	{{end}}
{{end}}
</ul>

</div>
{{if gt $t_idx 0}}
	<hr>
{{end}}
{{end}}
</div>
	<!-- <noscript> -->
{{template "front/create_reply_form_quick" .}}
	<!-- </noscript> -->
<hr>
{{template "front/post_list_nav_bot" .}}
{{/* remove return, add search box instead */}}
{{/* remove post/media count */}}
{{end}}

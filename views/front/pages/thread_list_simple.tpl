{{define "front/thread_list_simple"}}
<div class="board">
	{{range $t_idx, $thread := .thread_list}}
		{{$previewLen := (len $thread.Preview)}}
		{{if eq $previewLen 0}}
			{{continue}}
		{{end}}
		{{if gt $t_idx 0}}
			<hr>
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
					{{if eq $idx 0}}
						{{$pipe = (params "post" $post "type" $type "trp" $trp "subject" $thread.Subject)}}
					{{end}}
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
	{{else}}
		<div class="thread">
			<p>No threads yet, why don't you start one?</p>
		</div>
	{{end}}
</div>
{{end}}

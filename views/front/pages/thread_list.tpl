{{define "front/thread_list"}}
<ul>
{{range $thread := .thread_list}}
	<li>
		<div class="thread">
			<div class="postContainer opContainer">
				<div class="post op">
					<div class="postInfo desktop">
						<a href="/boards/{{$.board_code}}/threads/{{$thread.ID}}/">
							{{$thread.Content}}
						</a>
					</div>
				</div>
			</div>
		</div>
	</li>
{{else}}
	<li>No threads available!</li>
{{end}}
</ul>
{{end}}

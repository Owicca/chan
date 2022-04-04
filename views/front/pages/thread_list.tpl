{{define "front/thread_list"}}
<ul>
{{range $thread := .threads}}
	<li>
		<div class="thread">
			<div class="postContainer opContainer">
				<div class="post op">
					<div class="postInfo desktop">
						<a href="/threads/{{$thread.ID}}/">
						</a>
					</div>
					<blockquote class="postMessage">
						{{$thread.Primary.Content}}
					</blockquote>
				</div>
			</div>
		</div>
	</li>
{{else}}
	<li>No threads available!</li>
{{end}}
</ul>
{{end}}
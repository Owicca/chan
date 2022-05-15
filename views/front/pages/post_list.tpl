{{define "front/post_list"}}
<div class="board">
<div class="thread">
<ul>
{{range $idx, $post := .post_list}}
{{$type := "reply"}}
{{if eq $idx 0}}
	{{$type = "op"}}
{{end}}
	<li>
		<div class="postContainer {{$type}}Container">
			<div class="post {{$type}}">
				<div class="postInfo desktop">
					{{$post.ID}}
				</div>
				{{if $post.Media.Object_id}}
				<div class="file">
					<div class="fileText">
						File: 
						<a href="/static/media/{{$post.Media.Path}}" target="_blank">{{$post.Media.Name}}</a> 
						({{b2s $post.Media.Size}}, {{$post.Media.X}}x{{$post.Media.Y}})
					</div>
					<a class="fileThumb" href="/static/media/{{$post.Media.Path}}" target="_blank">
						<img src="/static/media/{{$post.Media.Thumb}}" alt="{{$post.Media.X}}" class="fileThumb--item" loading="lazy">
						<div class="mFileInfo mobile">{{b2s $post.Media.Size}}</div>
					</a>
				</div>
				{{end}}
				<blockquote class="postMessage">
					{{$post.Content}}
				</blockquote>
			</div>
		</div>
	</li>
{{end}}
</ul>
</div>
</div>
{{end}}

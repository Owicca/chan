{{define "front/post_one"}}
<li>
	<div class="postContainer {{.type}}Container">
		<div id="p{{.post.ID}}" class="post {{.type}}">
			<div class="postInfo">
				<span class="nameBlock">
					{{if .subject}}
						<span class="subject">{{.subject}}</span>
					{{end}}
					<span class="name">
					{{if .post.Name}}
						<span class="theName">{{.post.Name}}</span><span title="{{.trp}}" class="tripcode">{{.trp}}</span>
					{{else}}
						Anonymous
					{{end}}
					</span>
				</span>
				<span class="dateTime" data-utc="{{.post.Created_at}}">{{u2d .post.Created_at}}</span>
				<span class="postNum">
					<a href="#p{{.post.ID}}" title="Link to this post">No.</a>
					<a class="quotePost" data-id="{{.post.ID}}" title="Reply to this post">{{.post.ID}}</a>
				</span>
				<a href="#" class="postMenuBtn" title="Post menu" data-cmd="post-menu">▶</a>
				<div id="bl_{{.post.ID}}" class="backlink">
					<span>
						{{range $l := .post.LinkList}}
							<a href="#p{{$l.Dest}}" class="quotelink">&gt;&gt;{{$l.Dest}}</a>
						{{end}}
					</span>
				</div>
			</div>
			{{if .post.Media.Object_id}}
			<div class="file">
				<div class="fileText">
					File: 
					<a href="/static/media/{{.post.Media.Path}}" target="_blank">{{.post.Media.Name}}</a> 
					({{b2s .post.Media.Size}}, {{.post.Media.X}}x{{.post.Media.Y}})
				</div>
				<a class="fileThumb" href="/static/media/{{.post.Media.Path}}" target="_blank">
					<img src="/static/media/{{.post.Media.Thumb}}" alt="{{.post.Media.X}}" class="fileThumb--item" loading="lazy">
					<div class="mFileInfo mobile">{{b2s .post.Media.Size}}</div>
				</a>
			</div>
			{{end}}
			<blockquote class="postMessage">
				{{if eq .post.Deleted_at 0}}
					{{.post.Content}}
				{{else}}
					This post was deleted
				{{end}}
			</blockquote>
		</div>
	</div>
</li>
{{end}}

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
	<li id="p{{$post.ID}}">
		<div class="postContainer {{$type}}Container">
			<div class="post {{$type}}">
				<div class="postInfo">
					<span class="nameBlock">
						<span class="name">
						{{if $post.Name}}
							<span class="theName">{{$post.Name}}</span><span title="{{$trp}}" class="tripcode">{{$trp}}</span>
						{{else}}
							Anonymous
						{{end}}
						</span>
					</span>
					<span class="dateTime" data-utc="{{$post.Created_at}}">{{u2d $post.Created_at}}</span>
					<span class="postNum">
						<a href="#p{{$post.ID}}" title="Link to this post">No.</a>
						<a href="javascript:quote('{{$post.ID}}');" title="Reply to this post">{{$post.ID}}</a>
					</span>
					<a href="#" class="postMenuBtn" title="Post menu" data-cmd="post-menu">▶</a>
					<!-- <div id="bl_86931392" class="backlink"><span><a href="#p86932976" class="quotelink">&gt;&gt;86932976</a> </span></div> -->
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
	<!-- <noscript> -->
{{template "front/create_reply_form_quick" .}}
	<!-- </noscript> -->
</ul>

</div>
</div>
<hr>
{{template "front/post_list_nav_bot" .}}
{{end}}

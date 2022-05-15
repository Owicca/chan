{{define "front/post_list"}}
<div class="navLinks">
	[<a href="/boards/{{.board_code}}/" accesskey="a">Return</a>]
	[<a href="/boards/{{.board_code}}/catalog">Catalog</a>]
	[<a href="#bottom">Bottom</a>]
	<div class="thread-stats">
		<span class="ts-replies" data-tip="Replies">{{.stats.reply_count}}</span> /
		<span class="ts-images" data-tip="Images">{{.stats.media_count}}</span> /
		<!-- <span data-tip="Posters" class="ts-ips">76</span> / -->
		<span data-tip="Page" class="ts-page">{{.page_nr}}</span>
	</div>
</div>
<hr>
<div class="board">
<div class="thread">

<ul>
{{range $idx, $post := .post_list}}
{{$type := "reply"}}
{{if eq $idx 0}}
	{{$type = "op"}}
{{end}}
	<li id="p{{$post.ID}}">
		<div class="postContainer {{$type}}Container">
			<div class="post {{$type}}">
				<div class="postInfo">
					<span class="nameBlock"><span class="name">
					{{if $post.Name}}
						{{$post.Name}}
						{{if $post.SecureTripcode}}
							!!{{$post.SecureTripcode}}
						{{else if $post.Tripcode}}
							!{{$post.Tripcode}}
						{{end}}
					{{else}}
						Anonymous
					{{end}}
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
</ul>

</div>
</div>
<hr>
<div class="navLinks navLinksBot">
	<div class="open-qr-wrap">
		[<a href="#" class="open-qr-link">Post a Reply</a>]
	</div>
	[<a href="/boards/{{.board_code}}/" accesskey="a">Return</a>]
	[<a href="/boards/{{.board_code}}/catalog">Catalog</a>]
	[<a href="#top">Top</a>]
	<div class="thread-stats">
		<span class="ts-replies" data-tip="Replies">{{.stats.reply_count}}</span> /
		<span class="ts-images" data-tip="Images">{{.stats.media_count}}</span> /
		<!-- <span data-tip="Posters" class="ts-ips">76</span> / -->
		<span data-tip="Page" class="ts-page">{{.page_nr}}</span>
	</div>
</div>
{{end}}

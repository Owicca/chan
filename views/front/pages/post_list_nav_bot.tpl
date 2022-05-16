{{define "front/post_list_nav_bot"}}
<div class="navLinks navLinksBot">
	<div>
		[<a href="/boards/{{.board_code}}/" accesskey="a">Return</a>]
		[<a href="/boards/{{.board_code}}/catalog">Catalog</a>]
		[<a href="#top">Top</a>]
	</div>
	<div class="open-qr-wrap">
		[<a href="#" class="open-qr-link">Post a Reply</a>]
	</div>
	<div class="thread-stats">
		<span class="ts-replies" data-tip="Replies">{{.stats.reply_count}}</span> /
		<span class="ts-images" data-tip="Images">{{.stats.media_count}}</span> /
		<!-- <span data-tip="Posters" class="ts-ips">76</span> / -->
		<span data-tip="Page" class="ts-page">{{.page_nr}}</span>
	</div>
</div>
{{end}}

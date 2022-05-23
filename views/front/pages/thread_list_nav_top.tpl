{{define "front/thread_list_nav_top"}}
<div class="navLinks">
	<div>
		<input type="text" id="search-box" placeholder="Search OPs…">
		[<a href="/boards/{{.board_code}}/catalog">Catalog</a>]
	</div>
	<div class="open-qr-wrap">
		[<a href="#" class="open-qr-link">Post a Reply</a>]
	</div>
</div>
{{end}}

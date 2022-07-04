{{define "front/thread_list_nav_top"}}
<div class="navLinks">
	<div>
		<form action="/search/" method="post">
			<input name="search" type="text" id="search-box" placeholder="Search OPs…">
			<input name="board_code" type="hidden" value="{{.board_code}}">
			<input type="submit" value="Search">
		</form>
		<!--[<a href="/boards/{{.board_code}}/catalog">Catalog</a>]-->
	</div>
</div>
{{end}}

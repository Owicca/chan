{{define "front/pagination"}}
<div class="pagelist desktop">
	{{if gt .page 1 }}
	<div class="prev">
		<a href="/boards/{{$.board_code}}/{{dec .page}}/">Previous</a>
	</div>
	{{end}}
	<div class="pages">
		{{range $idx, $e := .page_helper}}
			[<a href="/boards/{{$.board_code}}/{{$e}}/">{{$e}}</a>]
		{{end}}
	</div>
	{{if lt .page (len .page_helper)}}
		<div class="next">
			<a href="/boards/{{$.board_code}}/{{inc .page}}/">Next</a>
		</div>
	{{end}}
	<div class="pages cataloglink">
		<a href="/boards/{{.board_code}}/catalog/">Catalog</a>
	</div>
</div>
{{end}}

{{define "front/search"}}
<h1>Search</h1>
<hr>
<form action="#" id="g-search-form" method="post">
	<input class="g-search-ctrl" id="js-sf-qf" name="search" type="text" placeholder="Search" {{if .search}}value="{{.search}}"{{end}} />
	<select class="g-search-ctrl" id="js-sf-bf" name="board_code">
		{{range $idx, $b := .board_list}}
			<option value="{{$b.Code}}" {{if eq $.board_code $b.Code}}selected{{end}}>/{{$b.Code}}/ - {{$b.Name}}</option>
		{{end}}
	</select>
	<button class="g-search-ctrl" id="js-sf-btn">Search</button>
</form>
<hr>
{{template "front/thread_list_simple" .}}
{{end}}

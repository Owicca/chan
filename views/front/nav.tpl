{{define "front/nav"}}
<div id="boardNavDesktop" class="desktop">
  <span class="boardList">
		[
		{{range $idx, $board := .navigation}}
			{{if gt $idx 0 }} / {{end}}
		<a href="/boards/{{$board.Code}}/" title="{{$board.Name}}">{{$board.Code}}</a>
		{{end}}
		]
		<span class="custom-menu-ctrl">
			[ <a data-cmd="custom-menu-edit" title="Edit Menu" href="#">Edit</a>]
		</span>
  </span>
  <span id="navtopright">
		[ <a href="javascript:void(0);" id="settingsWindowLink">Settings</a>]
		[ <a href="/search" title="Search">Search</a>]
		[ <a href="//p.4chan.org/" title="Mobile">Mobile</a>]
		[ <a href="/" target="_top">Home</a>]
	</span>
</div>
{{end}}

{{define "back/pagination"}}
<nav class="pagination-cnt">
  <ul class="pagination pagination-sm">
    {{if gt .page 1 }}
      <li class="page-item">
        <a class="page-link" href="/admin/threads/p{{dec .page}}/">Previous</a>
      </li>
    {{end}}
    {{range $idx, $e := .page_helper}}
      <li class="page-item {{if (eq $.page $e)}}active{{end}}">
        <a class="page-link" href="/admin/threads/p{{$e}}/">{{$e}}</a>
      </li>
    {{end}}
    {{if lt .page (len .page_helper)}}
      <li class="page-item">
        <a class="page-link" href="/admin/threads/p{{inc .page}}/">Next</a>
      </li>
    {{end}}
  </ul>
</nav>
{{end}}

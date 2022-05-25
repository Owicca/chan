{{define "back/thread_list"}}
<table class="table table-sm table-striped align-middle">
    <thead>
        <tr>
            <td scope="col">#</td>
            <td scope="col">Status</td>
            <td scope="col">Posts</td>
            <td scope="col">Actions</td>
        </tr>
    </thead>
    <tbody>
{{range $thread := .threads}}
    <tr>
        <td> 
            <a href="/admin/threads/{{$thread.ID}}/">
                {{$thread.ID}}
            </a>
        </td>
        <td> 
            {{if gt $thread.Deleted_at 0}}
                Deleted at: {{unixToUTC $thread.Deleted_at}}
            {{else}}
                Active
            {{end}}
        </td>
        <td> 
            <a href="/admin/threads/{{$thread.ID}}/posts/">
                {{len $thread.Preview}}
            </a>
        </td>
        <td>
            {{template "back/actions" params "update_name" "threads" "update_id" $thread.ID}}
        </td>
    </tr>
{{else}}
    <tr>
        <td colspan="4">No threads available!</td>
    </tr>
{{end}}
    </tbody>
</table>
{{end}}

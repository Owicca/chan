{{define "back/thread_list"}}

{{template "back/pagination" .}}

<table class="table table-sm table-striped align-middle">
    <thead>
        <tr>
            <td scope="col">#</td>
            <td scope="col">Board</td>
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
            <a href="/admin/boards/{{$thread.Board_id}}/">
                {{$thread.Board_id}}
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
        <td colspan="5">No threads available!</td>
    </tr>
{{end}}
    </tbody>
</table>

{{template "back/pagination" .}}

{{end}}

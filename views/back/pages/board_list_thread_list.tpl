{{define "back/board_list_thread_list"}}

{{template "back/pagination" .}}

<table class="table table-sm table-striped align-middle">
    <thead>
        <tr>
            <td scope="col">#</td>
            <td scope="col">Status</td>
            <td scope="col">Posts</td>
            <td scope="col">Thread</td>
        </tr>
    </thead>
    <tbody>
{{range $thread := .thread_list}}
    <tr>
        <td> 
            <a href="/admin/threads/{{$thread.ID}}/">
                {{$thread.ID}}
            </a>
        </td>
        <td> 
            {{if gt $thread.Deleted_at 0}}
            Deleted at {{$thread.Deleted_at}}
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
            <a href="/admin/threads/{{$thread.ID}}/">
                Update
            </a>
        </td>
    </tr>
{{else}}
    <tr>
        <td colspan="4">No threads available!</td>
    </tr>
{{end}}
    </tbody>
</table>

{{template "back/pagination" .}}

{{end}}

{{define "back/board_list_thread_list"}}
<table>
    <thead>
        <tr>
            <tr>
                ID
            </tr>
            <tr>
                Status
            </tr>
            <tr>
                Posts
            </tr>
            <tr>
                Thread
            </tr>
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
                Posts
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
        <td>No threads available!</td>
    </tr>
{{end}}
    </tbody>
</table>
{{end}}
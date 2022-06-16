{{define "back/post_list"}}
<table class="table table-sm table-striped align-middle">
    <thead>
        <tr>
            <td scope="col">#</td>
            <td scope="col">Thread</td>
            <td scope="col">Status</td>
            <td scope="col" width="50%">Content</td>
            <td scope="col">Actions</td>
        </tr>
    </thead>
    <tbody>
{{range $post := .posts}}
        <tr>
            <td>
                <a href="/admin/posts/{{$post.ID}}/">
                    {{$post.ID}}
                </a>
            </td>
            <td>
                <a href="/admin/threads/{{$post.Thread_id}}/">
                    {{$post.Thread_id}}
                </a>
            </td>
            <td>
                {{if gt $post.Deleted_at 0}}
                    Deleted
                {{else}}
                    {{range $label, $val := $.postStatusList}}
                        {{if eq $val $post.Status}}
                            {{$label}}
                        {{end}}
                    {{end}}
                {{end}}
            </td>
            <td>
                {{$post.Content}}
            </td>
            <td>
                <form method="post" action="/admin/posts/{{$post.ID}}/">
                    <input name="post_id" type="hidden" value="{{$post.ID}}" />
                    <input name="thread_id" type="hidden" value="{{$post.Thread_id}}" />
                    <input type="submit" value="Delete" />
                </form>
            </td>
        </tr>
{{else}}
        <tr>
            <td colspan="5">No posts found!</td>
        </tr>
{{end}}
    </tbody>
</table>
{{end}}

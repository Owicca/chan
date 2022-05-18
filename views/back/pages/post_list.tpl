{{define "back/post_list"}}
<table>
    <thead>
        <tr>
            <td>ID</td>
            <td>Status</td>
            <td width="50%">Content</td>
            <td>Thread</td>
            <td></td>
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
                <a href="/admin/threads/{{$post.Thread_id}}/">
                    {{$post.Thread_id}}
                </a>
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
            <p>No posts found!</p>
        </tr>
{{end}}
    </tbody>
</table>
{{end}}

{{define "back/post_list"}}
<table>
    <thead>
        <tr>
            <td>ID</td>
            <td>Status</td>
            <td width="50%" colspan="2">Content</td>
            <td>Is primary</td>
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
                {{range $label, $val := $.postStatusList}}
                    {{if eq $val $post.Status}}
                        {{$label}}
                    {{end}}
                {{end}}
            </td>
            <td>
                {{$post.Content}}
            </td>
            <td>
                {{if $post.Is_primary}}
                Primary
                {{else}}
                Not primary
                {{end}}
            </td>
            <td>
                <a href="/admin/threads/{{$post.Thread_id}}/">
                    {{$post.Thread_id}}
                </a>
            </td>
            <td>
                <a href="/admin/posts/{{$post.ID}}/">
                    Delete
                </a>
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
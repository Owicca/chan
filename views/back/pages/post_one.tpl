{{define "back/post_one"}}
<table>
	<thead>
		<th></th>
			<th>ID</th>
			<th>Created_at</th>
			<th>Deleted_at</th>
			<th>Tripcode</th>
			<th>SecureTripcode</th>
			<th>Status</th>
			<th>Thread_id</th>
			<th>Name</th>
			<th>Content</th>
			<th>Media</th>
	</thead>
	<tbody>
		<tr>
			<td>{{.post.ID}}</td>
			<td>{{.post.Created_at}}</td>
			<td>{{.post.Deleted_at}}</td>
			<td>{{.post.Tripcode}}</td>
			<td>{{.post.SecureTripcode}}</td>
			<td>{{.post.Status}}</td>
			<td>{{.post.Thread_id}}</td>
			<td>{{.post.Name}}</td>
			<td>{{.post.Content}}</td>
			<td>
				{{if .post.Media.Object_id}}
				<div class="file">
					<div class="fileText">
						File: 
						<a href="/static/media/{{.post.Media.Path}}" target="_blank">{{.post.Media.Name}}</a> 
						({{b2s .post.Media.Size}}, {{.post.Media.X}}x{{.post.Media.Y}})
					</div>
					<a class="fileThumb" href="/static/media/{{.post.Media.Path}}" target="_blank">
						<img src="/static/media/{{.post.Media.Thumb}}" alt="{{.post.Media.X}}" class="fileThumb--item" loading="lazy">
						<div class="mFileInfo mobile">{{b2s .post.Media.Size}}</div>
					</a>
				</div>
				{{else}}
					No media!
				{{end}}
			</td>
		</tr>
	</tbody>
</table>

<a href="/admin/threads/" class="btn btn-secondary">Back</a>
{{end}}

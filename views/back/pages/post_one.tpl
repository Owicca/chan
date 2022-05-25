{{define "back/post_one"}}
<table class="table table-sm table-striped align-middle text-start">
	<thead>
		<tr>
			<th scope="col">#</th>
			<th scope="col">Created at</th>
			<th scope="col">Deleted at</th>
			<th scope="col">Tripcode</th>
			<th scope="col">Secure Tripcode</th>
			<th scope="col">Status</th>
			<th scope="col">Thread</th>
			<th scope="col">Name</th>
			<th scope="col">Content</th>
			<th scope="col">Media</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>{{.post.ID}}</td>
			<td>{{unixToUTC .post.Created_at}}</td>
			<td>{{unixToUTC .post.Deleted_at}}</td>
			<td>{{.post.Tripcode}}</td>
			<td>{{.post.SecureTripcode}}</td>
			<td>{{.post.Status}}</td>
			<td>
				<a href="/admin/threads/{{.post.Thread_id}}/">
					{{.post.Thread_id}}
				</a>
			</td>
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

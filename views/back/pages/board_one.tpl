{{define "back/board_one"}}
{{if .board}}
<form method="POST" action="/admin/boards/{{.board.ID}}/">
	<input type="hidden" id="board_id" name="board_id" value="{{.board.ID}}" />
	<div class="input-group input-group-sm mb-3">
		<label for="name" class="input-group-text">Name: </label>
		<input type="text" id="name" class="form-control" name="name" value="{{.board.Name}}" />
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="code" class="input-group-text">Code: </label>
		<input type="text" id="code" class="form-control" name="code" value="{{.board.Code}}" />
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="description" class="input-group-text">Description: </label>
		<textarea id="description" class="form-control" name="description">{{.board.Description}}</textarea>
	</div>
	<div class="actions">
		<input type="submit" class="btn btn-success" value="Submit" />
		<a href="/admin/boards/" class="btn btn-secondary">Back</a>
	</div>
</form>
{{else}}
	<p>board not found</p>
{{end}}
{{end}}
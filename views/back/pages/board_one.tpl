{{define "back/board_one"}}
{{if .board}}
<form method="POST" action="/admin/boards/{{if .board.ID}}{{.board.ID}}/{{end}}">
    {{if .board.ID}}
    <input type="hidden" id="id" name="id" value="{{.board.ID}}" />
    {{end}}
	<div class="input-group input-group-sm mb-3">
		<label for="name" class="input-group-text">Name: </label>
		<input type="text" id="name" class="form-control" name="name" value="{{.board.Name}}" />
	</div>
	<div class="input-group input-group-sm mb-3">
		<label for="code" class="input-group-text">Code: </label>
		<input type="text" id="code" class="form-control" name="code" value="{{.board.Code}}" />
	</div>
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Topic: </label>
        <select class="form-select d-flex justify-content-start me-3" name="topic_id">
    {{range $topic := .topic_list}}
            <option value="{{$topic.ID}}" {{if eq $topic.ID $.board.Topic_id}}selected="selected"{{end}}>{{$topic.Name}}</option>
    {{else}}
        <p>No topics found!</p>
    {{end}}
        </select>
    </div>
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Status: </label>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="0" {{if eq .board.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Active</label>
        </div>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="1" {{if gt .board.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Deleted</label>
        </div>
        {{if gt .board.Deleted_at 0}}
        <p>Deleted at: {{unixToUTC .board.Deleted_at}}</p>
        {{end}}
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
	<p>Board not found</p>
{{end}}
{{end}}
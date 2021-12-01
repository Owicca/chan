{{define "back/thread_one"}}
{{if .thread}}
<h1>{{.thread.ID}}</h1>
<form method="POST" action="/admin/threads/{{.thread.ID}}/">
    <input type="hidden" id="thread_id" name="thread_id" value="{{.thread.ID}}" />
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Status: </label>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="0" {{if eq .thread.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Active</label>
        </div>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="1" {{if gt .thread.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Deleted</label>
        </div>
        {{if gt .thread.Deleted_at 0}}
        <p>Deleted at: {{.thread.Deleted_at}}</p>
        {{end}}
    </div>
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Board: </label>
        <select class="form-select d-flex justify-content-start me-3" name="board_id">
    {{range $board := .boardList}}
            <option value="{{$board.ID}}" {{if eq $board.ID $.thread.Board_id}}checked="checked"{{end}}>{{$board.Name}}</option>
    {{else}}
        <p>No boards found!</p>
    {{end}}
        </select>
    </div>
    <div class="actions">
        <input type="submit" class="btn btn-success" value="Submit" />
        <a href="/admin/threads/" class="btn btn-secondary">Back</a>
    </div>
</form>
{{else}}
    <p>Thread not found</p>
{{end}}
{{end}}
{{define "back/topic_one"}}
{{if .topic}}
<form method="POST" action="/admin/topics/{{if .topic.ID}}{{.topic.ID}}/{{end}}">
    {{if .topic.ID}}
    <input type="hidden" id="id" name="id" value="{{.topic.ID}}" />
    {{end}}
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Name: </label>
        <input class="form-control" type="text" name="name" value="{{.topic.Name}}" />
    </div>
    <div class="input-group input-group-sm mb-3">
        <label class="input-group-text">Status: </label>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="0" {{if eq .topic.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Active</label>
        </div>
        <div class="form-check d-flex justify-content-start me-3">
            <input type="radio" id="deleted_at" class="form-check-input me-1" name="deleted_at" value="1" {{if gt .topic.Deleted_at 0}}checked="checked"{{end}} />
            <label for="deleted_at" class="form-check-label">Deleted</label>
        </div>
        {{if gt .topic.Deleted_at 0}}
        <p>Deleted at: {{unixToUTC .topic.Deleted_at}}</p>
        {{end}}
    </div>
    <div class="actions">
        <input type="submit" class="btn btn-success" value="Submit" />
        <a href="/admin/topics/" class="btn btn-secondary">Back</a>
    </div>
</form>
{{else}}
    <p>Topic not found</p>
{{end}}
{{end}}
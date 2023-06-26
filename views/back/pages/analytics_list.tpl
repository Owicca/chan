{{define "back/analytics_list"}}

<div class="container">
	<ul class="nav nav-tabs flex-column flex-sm-row" id="tabs" role="tablist">
		{{range $tab := .types}}
			<li class="nav-item flex-sm-fill" role="presentation">
				<button id="{{$tab.ID}}-tab" class="nav-link {{if $tab.IsActive}}active{{else}}{{end}}" data-bs-toggle="tab" data-bs-target="#{{$tab.ID}}-tab-pane" type="button">{{$tab.Name}}</button>
			</li>
		{{end}}
	</ul>

	<div class="tab-content" id="tabContent">
		{{range $tab := .types}}
			<div class="tab-pane fade {{if $tab.IsActive}}show active{{else}}{{end}}" id="{{$tab.ID}}-tab-pane" role="tabpanel" aria-labelledby="{{$tab.ID}}-tab" tabindex="0">{{$tab.Name}}</div>
		{{end}}
	</div>
</div>

{{end}}

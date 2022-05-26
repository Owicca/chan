{{define "back/alerts"}}
<div id="alerts">
	{{range $idx, $alert := .data.flash_list_info}}
		<div class="toast bg-info toast-{{$idx}}-info d-flex justify-content-between align-items-center" role="alert" aria-live="assertive" aria-atomic="true">
			<div class="toast-body">
				{{$alert}}
			</div>
			<button type="button" class="btn-close" data-bs-dissmis="toast-{{$idx}}-info" aria-label="close"></button>
		</div>
	{{end}}
	{{range $idx, $alert := .data.flash_list_success}}
		<div class="toast bg-success toast-{{$idx}}-success d-flex justify-content-between align-items-center" role="alert" aria-live="assertive" aria-atomic="true">
			<div class="toast-body">
	 			{{$alert}}
			</div>
			<button type="button" class="btn-close" data-bs-dissmis="toast-{{$idx}}-success" aria-label="close"></button>
		</div>
	{{end}}
	{{range $idx, $alert := .data.flash_list_error}}
		<div class="toast bg-danger toast-{{$idx}}-error d-flex justify-content-between align-items-center" role="alert" aria-live="assertive" aria-atomic="true">
			<div class="toast-body">
				{{$alert}}
			</div>
			<button type="button" class="btn-close" data-bs-dissmis="toast-{{$idx}}-error" aria-label="close"></button>
		</div>
	{{end}}
</div>
{{end}}

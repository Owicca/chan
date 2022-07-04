{{define "front/thread_list"}}
{{$form_action := (printf "/boards/%s/" .board_code)}}
{{$form_params := (params "form_action" $form_action "form_button_label" "Start a New Thread" "create_thread" true "errors" .errors)}}
{{template "front/create_reply_form" $form_params}}
<hr>
{{template "front/thread_list_nav_top" .}}
{{/* remove return, add search box instead */}}
{{/* remove post/media count */}}
<hr>
{{template "front/thread_list_simple" .}}
	<!-- <noscript> -->
{{template "front/create_reply_form_quick" .}}
	<!-- </noscript> -->
<hr>
{{template "front/pagination" .}}
{{end}}

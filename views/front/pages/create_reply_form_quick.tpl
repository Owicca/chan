{{define "front/create_reply_form_quick"}}
<form id="quickReply" class="hidden" action="/boards/{{.board_code}}/threads/{{.thread_id}}/" method="post" enctype="multipart/form-data">
	<div id="qrForm">
		<div>
			<input name="name" type="text" tabindex="0" placeholder="Name">
		</div>
		<div>
			<textarea name="content" cols="48" rows="4" tabindex="0" placeholder="Comment" wrap="soft"></textarea>
		</div>
		<div>
			<input id="qrFile" name="media" type="file" tabindex="0" size="19" title="Shift + Click to remove the file">
			<input type="submit" value="Post" tabindex="0">
		</div>
	</div>
</form>
{{end}}

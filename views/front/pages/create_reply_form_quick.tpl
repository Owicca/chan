{{define "front/create_reply_form_quick"}}
<div id="quickReply" class="extPanel reply hidden">
	<div id="qrHeader" class="drag postblock">
		<span>Reply to Thread No.</span>
		<span id="qrTid"></span>
		<img alt="X" src="/static/media/image/cross.png" id="qrClose" class="extButton" title="Close Window">
	</div>
	<form id="quickReplyForm" action="{{.form_action}}" method="POST" enctype="multipart/form-data" data-board-code="{{.board_code}}">
		<div id="qrForm">
			<div>
				<input class="name" name="name" type="text" tabindex="1" placeholder="Name">
				{{template "front/errors" (.errors.Get "name")}}
			</div>
			<div>
				<input name="email" type="text" tabindex="0" id="qrEmail" placeholder="Options">
			</div>
			<div id="linkCnt" class="backlink">
			</div>
			<div>
				<textarea name="content" cols="48" rows="4" tabindex="0" placeholder="Content" wrap="soft"></textarea>
				{{template "front/errors" (.errors.Get "content")}}
			</div>
			<div id="qrCaptchaContainer" class="t-qr-root"></div>
			<div>
				<input id="postFile" name="media" type="file" tabindex="0" size="19" title="Shift + Click to remove the file">
				{{template "front/errors" (.errors.Get "media")}}
				<input type="submit" value="{{.form_button_label}}" tabindex="0">
			</div>
		</div>
	</form>
	<div id="qrError">
		{{template "front/errors" (.errors.Get "misc")}}
	</div>
</div>
{{end}}

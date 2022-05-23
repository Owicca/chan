{{define "front/create_reply_form"}}
<form id="reply" action="{{.form_action}}" method="post" enctype="multipart/form-data">
	<div id="togglePostFormLink">[<a href="#">{{.form_button_label}}</a>]</div>
	<table class="postForm hidden" id="postForm">
		<tbody>
			<tr data-type="Name">
				<td>Name</td>
				<td>
					<input class="name" name="name" type="text" tabindex="1" placeholder="Anonymous">
					<input type="submit" value="Post" tabindex="6">
				</td>
			</tr>
			{{if .create_thread}}
				<tr data-type="Subject">
					<td>Subject</td>
					<td>
						<input name="subject" type="text" tabindex="3">
					</td>
				</tr>
			{{end}}
			<tr data-type="Content">
				<td>Content</td>
				<td>
					<textarea name="content" cols="48" rows="4" tabindex="4" wrap="soft"></textarea>
				</td>
			</tr>
			<tr data-type="File">
				<td>File</td>
				<td>
					<input id="postFile" name="media" type="file" tabindex="7" {{if .create_thread}}required="required"{{end}}>
				</td>
			</tr>
			<tr class="rules">
				<td colspan="2">
					<ul class="rules">
						<li>Please read the <a href="/rules#g/">Rules</a> and <a href="/faq/">FAQ</a> before posting.</li>
						<li>You may highlight syntax and preserve whitespace by using [code] tags.</li>
						<li>There are <strong id="unique-ips">#@#</strong> posters in this thread.</li>
					</ul>
				</td>
			</tr>
		</tbody>
		<tfoot>
			<tr>
				<td colspan="2">
					<div id="postFormError"></div>
				</td>
			</tr>
		</tfoot>
	</table>
</form>
{{end}}

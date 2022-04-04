{{define "front/index"}}
<div id="doc">
	<header id="hd">
		<div id="logo-fp"></div>
	</header>
	<main id="bd">
		<div class="box-outer" id="announce">
			<div class="box-inner">
				<div class="boxbar">
					<h2>What is {{.site.name}}?</h2>
					<a data-cmd="x-wot" href="#" class="closebutton"></a>
				</div>
				<div class="boxcontent">
					<div id="wot-cnt">
						{{.site.welcome}}
					</div>
				</div>
			</div>
		</div>
		<div id="boards" class="box-outer top-box">
			<div class="box-inner">
				<div class="boxbar">
					<h2>Boards</h2>
				</div>
				<div class="boxcontent">
				{{range $topic := .topic_list}}
					<div class="column">
						<h3 class="col">{{$topic.Name}}</h3>
						<ul>
						{{range $board := $topic.BoardList}}
							<li class="row">
								<div class="card">
									{{range $media := $board.MediaList}}
									<img class="image rounded card-img-top" src="{{$media.Path}}">
									{{end}}
									<div class="card-body">
										<h5 class="card-title">{{$board.Name}}</h5>
										<p class="card-text">{{$board.Description}}</p>
										<a href="/boards/{{$board.Code}}/" class="btn btn-primary">View</a>
									</div>
								</div>
							</li>
						{{else}}
							<li>No boards available!</li>
						{{end}}
						</ul>
					</div>
				{{end}}
				</div>
			</div>
		</div>
	</main>
	<footer id="ft">
		<ul>
			<li class="fill"></li>
			<li class="first"><a href="/">Home</a></li>
		</ul>
	</footer>
</div>
{{end}}
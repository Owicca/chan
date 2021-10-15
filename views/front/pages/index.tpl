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
				{{range $col, $boardList := .topics}}
					<div class="column">
						<h3 class="col">{{$col}}</h3>
						<ul>
						{{range $idx, $board := $boardList}}
							<li>
								<a href="{{$board.Path}}">{{$board.Name}}</a>
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
package templates

const HtmlFullTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SPM-Go Results</title>
    <style>
        body, html {
            margin:0; padding: 0;
            height: 100%;
        }
        body {
			padding: 10px;
            font-family: Helvetica Neue, Helvetica, Arial;
            font-size: 14px;
            color:#333;
        }
        .analysis-summary {
            border-collapse: collapse;
            width: 100%;
        }
        .analysis-summary tr { border-bottom: 1px solid #bbb; }
        .analysis-summary td, .analysis-summary th { padding: 10px; }
        .analysis-summary tbody { border: 1px solid #bbb; }
        .analysis-summary td { border-right: 1px solid #bbb; background: #fff4c2; }
        .analysis-summary th {
            border: 1px solid #bbb;
            text-align: center;
            font-weight: bold;
            white-space: nowrap;
        }
        .green > td { background: rgb(230,245,208) !important; }
        .red > td { background: #FCE1E5 !important; }
        .bold { font-weight: bold; }
        .right { text-align: right; }
        .right-spaced { padding-right: 40px !important; }
		.footer {
			color: rgba(0,0,0,0.5);
			text-align: center;
			margin-top: 10px;
			height: 50px;
		}
		.files {
			min-width: 200px;
		}
		.dependencies {
			min-width: 400px;
			max-width: 600px;
		}
		.dependants {
			min-width: 400px;
			max-width: 600px;
		}
		.abstractness {
			min-width: 200px;
		}
		.box {
			display: inline-block;
			padding: 10px;
			vertical-align: top;
			border: 2px solid #aaa;
			border-radius: 10px;
		}
		.box-title {
			width: 100%;
			text-align: center;
			font-weight: bold;
		}
		ul {
			margin-left: -20px;
		}
		.package-details {
		}
		.package-details > td {
			padding-left: 40px;
			background-color: #fcfcfc;
			border-bottom: 1px solid #bbb;
		}
    </style>
</head>
<body>
<h1>Software Package Metrics for Go - v{{.Version}}</h1>
<h2>Full analysis for Module: {{.Module}}</h2>
<div class="results">
<table class="analysis-summary">
    <thead>
    <tr>
        <th style="width:120px;" rowspan="2">Name</th>
        <th rowspan="2">Package Path</th>
        <th rowspan="2" style="width:60px;">Files</th>
        <th colspan="3">Instability</th>
        <th colspan="3">Abstractness</th>
        <th colspan="3">Distance</th>
    </tr>
    <tr>
        <th style="width:120px;">C<sub>Afferent</sub></th>
        <th style="width:120px;">C<sub>Efferent</sub></th>
        <th style="width:120px;">Instability</th>
        <th style="width:120px;">N<sub>Abstractions</sub></th>
        <th style="width:120px;">N<sub>Implementations</sub></th>
        <th style="width:120px;">Abstractness</th>
        <th style="width:120px;">Distance</th>
    </tr>
    </thead>
    <tbody>
    {{ range .Packages }}
    <tr class="{{if le .Distance 0.1}}green{{else if gt .Distance 0.6}}red{{end}}" onClick="javasript:toggleDetails('{{.Path}}-details')">
        <td class="">{{ .Name }}</td>
        <td class="">{{ .Path }}</td>
        <td class="right">{{ .FilesCount }}</td>
        <td class="right right-spaced">{{ .AfferentCoupling }}</td>
        <td class="right right-spaced">{{ .EfferentCoupling }}</td>
        <td class="bold right right-spaced">{{ .Instability }}</td>
        <td class="right right-spaced">{{ .AbstractionsCount }}</td>
        <td class="right right-spaced">{{ .ImplementationsCount }}</td>
        <td class="bold right right-spaced">{{ .Abstractness }}</td>
        <td class="bold right right-spaced">{{ .Distance }}</td>
    </tr>
	<tr id="{{.Path}}-details" class="package-details" style="display:none;">
		<td colspan="10">
			<div style="margin-bottom: 15px;"><b>Package details for: </b>{{ .Path }}</div>
			<div class="box files">
				<div class="box-title">Files</div>
				<ul>
					{{ range .Files }}
					<li>{{.}}</li>
					{{ end }}
				</ul>
			</div>
			{{ if .Dependencies }}
			<div class="box dependencies">
				<div class="box-title">Dependencies</div>
				<ul>
					<li>
						<b>Internals</b>
						<ul>
							{{ range .Dependencies.Internals }}
							<li>{{.}}</li>
							{{ end }}
						</ul>
					</li>
					<li>
						<b>Externals</b>
						<ul>
							{{ range .Dependencies.Externals }}
							<li>{{.}}</li>
							{{ end }}
						</ul>
					</li>
					<li>
						<b>Standard</b>
						<ul>
							{{ range .Dependencies.Standard }}
							<li>{{.}}</li>
							{{ end }}
						</ul>
					</li>
				</ul>
			</div>
			{{ end }}
			{{ if .Dependants }}
			<div class="box dependants">
				<div class="box-title">Dependants</div>
				<ul>
					{{ range .Dependants }}
					<li>{{.}}</li>
					{{ end }}
				</ul>
			</div>
			{{ end }}
			{{ if .AbstractnessDetails }}
			<div class="box abstractness">
				<div class="box-title">Abstractness Details</div>
				<ul>
					<li><b>Methods: </b>{{ .AbstractnessDetails.MethodsCount }}</li>
					<li><b>Functions: </b>{{ .AbstractnessDetails.FunctionsCount }}</li>
					<li><b>Interfaces: </b>{{ .AbstractnessDetails.InterfacesCount }}</li>
					<li><b>Structs: </b>{{ .AbstractnessDetails.StructsCount }}</li>
				</ul>
			</div>
			{{ end }}
		</td>
	</tr>
    {{ end }}
    </tbody>
</table>
</div>
<div class="footer">
	Report generated by
	<a href="https://pkg.go.dev/github.com/fdaines/spm-go" target="_blank">SPM-Go</a>
	at {{.TimeStamp}}
</div>
<script language="javascript">
function toggleDetails(element) {
	var target = document.getElementById(element)
	if(target.style.display == "none") {
		document.getElementById(element).style.display = "contents";
	} else {
		document.getElementById(element).style.display = "none";
	}
}
</script>
</body>
</html>`
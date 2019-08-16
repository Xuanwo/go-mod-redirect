package handler

import (
	"html/template"
)

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
<h1>{{.Host}}</h1>
<ul>
{{range .Paths}}<li><a href="https://godoc.org/{{.}}">{{.}}</a></li>{{end}}
</ul>
</html>
`))

var importTmpl = template.Must(template.New("import").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Import}} {{.VCS}} {{.Repo}}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.Import}}/{{.Subpath}}">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/{{.Import}}/{{.Subpath}}">see the package on godoc</a>.
</body>
</html>`))

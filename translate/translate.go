// translate.go has handlers for displaying the actual phonetic alphabet translations.
package translate

import (
	"fmt"
	"html/template"
	"net/http"
)

var mainString = `<!DOCTYPE html>
<html lang="en">
<head>
	<title>Phonetic Alphabet{{if .Phrase}}: {{.Phrase}}{{end}}</title>
    <link rel="shortcut icon" href="/s/images/chart-32.png" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">
	<script>
		(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
		m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
		ga('create', 'UA-20919137-2', 'nato-alpha.appspot.com');
		ga('send', 'pageview');
	</script>
</head>
<body>
	<div class="container">
		<h1>Phonetic Alphabet Translator</h1>
		{{if .Phrase}}
			<p class="lead">
			<b>{{.Phrase}}</b> is:
			{{range .Translation}}
				<span title="{{.Pronunciation}}">{{.Name}}</span>
			{{end}}
			</p>
		{{end}}

		<form method="get" action="/" class="form-inline" role="form">
			<div class="form-group">
				<input type="text" name="phrase" placeholder="Enter a phrase" value="{{.Phrase}}" class="form-control">
			</div>
			<input type="submit" value="Translate" class="btn btn-default">
		</form>
		<hr>
		<p><small>
			Using <b>{{.AlphabetName}}</b> phonetic alphabet.<br>
			<a href="http://www.michaelkelly.org">Michael Kelly</a>.
			<a href="https://github.com/mjkelly/nato-alpha">Code is on GitHub</a>.
			Data is public domain.
		</small></p>
	</div>
</body>
</html>
`

var mainTmpl = template.Must(template.New("main").Parse(mainString))
var alpha = MustLoad("./data/faa.csv", "FAA")

type translationPage struct {
	Phrase       string
	AlphabetName string
	Translation  []*Translation
}

func init() {
	http.HandleFunc("/", handler)
}

func writeError(w http.ResponseWriter, err error) {
	fmt.Sprintln(w, err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		writeError(w, err)
		return
	}
	// The form puts the phrase in an actual form value, but we want to support
	// short URLs, so we check both.
	var phrase string
	phrase_form := r.Form["phrase"]
	if r.URL.Path != "/" {
		phrase = r.URL.Path[1:]
	} else if len(phrase_form) > 0 {
		phrase = phrase_form[0]
	}

	d := translationPage{
		Phrase:       phrase,
		AlphabetName: alpha.Type,
		Translation:  alpha.Translate(phrase),
	}
	mainTmpl.Execute(w, d)
	if err != nil {
		writeError(w, err)
	}
}

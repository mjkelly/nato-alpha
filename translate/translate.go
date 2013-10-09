// translate.go has handlers for displaying the actual phonetic alphabet translations.
package translate

import (
	"fmt"
	"html/template"
	"net/http"
)

var mainString = `
<!DOCTYPE html>
<html lang="en">
<head>
	<title>Phonetic Alphabet: {{.Phrase}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">
</head>
<body>
	<div class="container">
		<h1>Phonetic Alphabet Translator</h1>
		<p>
		{{if .Phrase}}
			<b>{{.Phrase}}</b> is:
			{{range .Translation}}
				<span title="{{.Pronunciation}}">{{.Name}}</span>
			{{end}}
		{{end}}
		</p>

		<hr>
		<form method="get" action="/" class="form-inline" role="form">
			<div class="form-group">
				<input type="text" name="phrase" value="{{.Phrase}}" class="form-control">
			</div>
			<input type="submit" value="Translate" class="btn btn-default">
		</form>
		<p><small>Using <b>{{.AlphabetName}}</b> phonetic alphabet.</small></p>
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
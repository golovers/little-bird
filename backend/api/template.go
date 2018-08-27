package api

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/golovers/xtract"
)

var (
	indexTmpl        = parseTemplate("index.html")
	postTmpl         = parseTemplate("post.html")
	postTrendingTmpl = parseTemplate("post-trending.html")
	postAddTmpl      = parseTemplate("post-add.html")
	postMineTmpl     = parseTemplate("post-mine.html")
)

// parseTemplate applies a given file to the body of the base template.
func parseTemplate(filename string) *appTemplate {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))
	fn := template.FuncMap{
		"htmlNoEscape": htmlNoEscape,
		"htmlShort":    htmlShort,
	}
	tmpl.Funcs(fn)
	// Put the named file into a template called "body"
	path := filepath.Join("templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %v", err))
	}
	template.Must(tmpl.New("body").Parse(string(b)))

	return &appTemplate{tmpl.Lookup("base.html")}
}

// appTemplate is a user login-aware wrapper for a html/template.
type appTemplate struct {
	t *template.Template
}

// Execute writes the template using the provided data, adding login and user
// information to the base template.
func (tmpl *appTemplate) Execute(w http.ResponseWriter, r *http.Request, data interface{}) *appError {
	d := struct {
		Data        interface{}
		AuthEnabled bool
		Profile     *Profile
		LoginURL    string
		LogoutURL   string
	}{
		Data:        data,
		AuthEnabled: oauthConfig != nil,
		LoginURL:    "/login?redirect=" + r.URL.RequestURI(),
		LogoutURL:   "/logout?redirect=" + r.URL.RequestURI(),
	}

	if d.AuthEnabled {
		// Ignore any errors.
		d.Profile = profileFromSession(r)
	}

	if err := tmpl.t.Execute(w, d); err != nil {
		return appErrorf(err, "could not write template: %v", err)
	}
	return nil
}

func htmlNoEscape(v string) template.HTML {
	return template.HTML(v)
}

func htmlShort(v string, n int) string {
	return xtract.ValueLim(v, n)
}

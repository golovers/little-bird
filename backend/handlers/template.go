package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

var (
	articleListSubTemplate = templateItem{"articleList", "articles/list/list.html"}
	indexTmpl              = parseTemplateWithSubTemplates("index.html", articleListSubTemplate)
	articleDetailsTmpl     = parseTemplate("articles/details/details.html")
	trendingArticlesTmpl   = parseTemplateWithSubTemplates("articles/trending/trending.html", articleListSubTemplate)
	newArticleTmpl         = parseTemplate("articles/new/new.html")
	myArticlesTmpl         = parseTemplateWithSubTemplates("articles/mine/mine.html", articleListSubTemplate)
	editArticleTmpl        = parseTemplate("articles/edit/edit.html")
	notFoundTmpl           = parseTemplate("notfound.html")
)

type templateItem struct {
	name     string
	filename string
}

// parseTemplate applies a given file to the body of the base template.
func parseTemplate(filename string) *appTemplate {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))
	fn := template.FuncMap{
		"htmlNoEscape": htmlNoEscape,
		"shortDate":    shortDate,
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

func parseTemplateWithSubTemplates(filename string, templates ...templateItem) *appTemplate {
	tmpl := parseTemplate(filename)
	for _, item := range templates {
		path := filepath.Join("templates", item.filename)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("could not read template: %v", err))
		}
		template.Must(tmpl.t.New(item.name).Parse(string(b)))
	}

	return &appTemplate{tmpl.t.Lookup("base.html")}
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

func shortDate(d time.Time) string {
	t := time.Since(d)
	if t.Seconds() < 60 {
		return fmt.Sprintf("%d seconds ago", int(t.Seconds()))
	}
	since := t.Hours()
	if since < 1 {
		return fmt.Sprintf("%d minutes ago", int(t.Minutes()))
	}
	if since < 24 {
		return fmt.Sprintf("%d hours ago", int(since))
	}
	if since < 7*24 {
		return fmt.Sprintf("%d days ago", int(since)/(24))
	}
	if since < 7*24*4 {
		return fmt.Sprintf("%d weeks ago", int(since)/(7*24))
	}
	if since < 7*24*4*12 {
		return fmt.Sprintf("%d months ago", int(since)/(7*24*4))
	}
	return fmt.Sprintf("%d years ago", int(since)/(7*24*4*12))
}

package view

import (
	"html/template"
	"io"
	"net/http"
)

type View struct {
	templ *template.Template
}

var view *View

func NewView() *View {
	if view == nil {
		v := &View{
			templ: template.Must(template.ParseGlob( "pkg/templates/*.html")),
		}
		view = v
	}
	return view
}

func (v *View) Static(w io.Writer, name string) error {
	return v.templ.ExecuteTemplate(w, name, nil)
}

func (v *View) Render(w io.Writer, name string, data interface{}) error {
	return v.templ.ExecuteTemplate(w, name, data)
}

func (v *View) NotFound(w io.Writer, req *http.Request) error {
	return v.templ.ExecuteTemplate(w, "404.html", dataNotFound(req))
}

// Data groups global data that is dynamic on almost every HTML page.
type Data struct {
	Title string
	Header string
	Message string
	Data map[string]interface{}
}

var dataNotFound = func(req *http.Request) Data {
	return Data{
		Title:   "404 - Not Found",
		Header:  "Sorry!",
		Message: "The requested resource could not be found",
		Data: map[string]interface{}{
			"method": req.Method,
			"path": req.URL.Path,
		},
	}
}
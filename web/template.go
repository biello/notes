package web

import "html/template"

func init() {

}

func GetTemplate(filename string) (*template.Template, error) {
	return template.New(filename).ParseFiles(getTemplatePath() + filename)
}

func getTemplatePath() string {
	return "./web/templates/"
}

var EmptyPageText = []byte(`# Empty page
So this is an empty page`)

var Login = template.Must(GetTemplate("login.gohtml"))

var Show = template.Must(GetTemplate("show.gohtml"))

var Edit = template.Must(GetTemplate("edit.gohtml"))

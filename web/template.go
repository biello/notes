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

var UnauthorizedText = []byte(`# Unauthorized
Current account cannot access this page`)

var Login = template.Must(GetTemplate("login.gohtml"))

var Password = template.Must(GetTemplate("password.gohtml"))

var Register = template.Must(GetTemplate("register.gohtml"))

var Show = template.Must(GetTemplate("show.gohtml"))

var Notes = template.Must(GetTemplate("notes.gohtml"))

var Edit = template.Must(GetTemplate("edit.gohtml"))

package views

import (
	"html/template"
	"log"
	"path/filepath"
)

const (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(
		files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings representing
// the layout files used in our app
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

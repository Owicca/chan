package infra

import (
	"net/http"
	"bytes"
	"log"
	"strings"
	"fmt"
	"errors"

	stdtemplate "html/template"
	customtemplate "html/template"

	// customtemplate "github.com/alecthomas/template"
	// blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Template struct {
	templates *stdtemplate.Template
	custom 	  *stdtemplate.Template
	funcMap   stdtemplate.FuncMap
}

func NewTemplate() *Template {
	funcMap := customtemplate.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"last": func(a []int) int {
			if len(a) == 0 {
				return -1
			}
			return a[len(a)-1]
		},
		"asHTML": func(html string) customtemplate.HTML {
			return customtemplate.HTML(html)
		},
		"params": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("'params' should be called with pairs of values")
			}

			dict := make(map[string]interface{}, len(values))
			for i := 0; i < len(values); i+=2 {
				k, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("%d th key is not a string", i/2)
				}
				dict[k] = values[i+1]
			}

			return dict, nil
		},
	}

	templatePagePath := []string{
		"views/front/*.tpl",
		"views/back/*.tpl",
	}
	pagesPath := []string{
		"views/front/pages/*.tpl",
		"views/back/pages/*.tpl",
	}
	var (
		templates = customtemplate.New("template")
		tplPages = customtemplate.New("page")
	)

	for _, pathGlob := range templatePagePath {
		templates = customtemplate.Must(templates.New("template").Funcs(funcMap).ParseGlob(pathGlob))
	}
	for _, pathGlob := range pagesPath {
		tplPages = customtemplate.Must(tplPages.New("page").Funcs(funcMap).ParseGlob(pathGlob))
	}

	return &Template{
		templates: templates,
		custom: tplPages,
		funcMap: funcMap,
	}
}

func (t *Template) JSEscapeString(s string) string {
	return customtemplate.JSEscapeString(s)
}

func (t *Template) Render(w http.ResponseWriter, status int, name string, data interface{}) error {
	w.WriteHeader(status)

	buffer := bytes.NewBufferString("")
	t.custom.ExecuteTemplate(buffer, name, data)

	content := map[string]interface{}{
		"page": buffer.String(),
	}

	baseTplName := "template"
	environment := "front"
	if strings.HasPrefix(name, "back") {
		environment = "back"
	}
	baseTplName = strings.Join([]string{environment, baseTplName}, "/")

	err := t.templates.ExecuteTemplate(w, baseTplName, content)
	if err != nil {
		log.Fatalf("Could not render %s (%s)", name, err)
	}

	return err
}

func (t *Template) StringToHTML(s string) stdtemplate.HTML {
	return stdtemplate.HTML(s)
}

// func (t *Template) MarkdownToHTML(s string) stdtemplate.HTML {
// 	return stdtemplate.HTML(blackfriday.Run([]byte(s)))
// }
package html

import (
	"html/template"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin/render"
)

type (
	HtmlTemplMap map[string][]string
	HtmlFuncMap  map[string]template.FuncMap
)

type viewer struct {
	box    *rice.Box
	templs HtmlTemplMap
	funcs  HtmlFuncMap
}

// Instance supply render string
func (v *viewer) Instance(name string, data interface{}) render.Render {
	templ := template.New(name)

	if htmls, ok := v.templs[name]; ok {
		var builder strings.Builder

		for _, name := range htmls {
			builder.WriteString(v.box.MustString(name))
		}

		templ = template.Must(templ.Parse(builder.String()))
	}

	if funcs, ok := v.funcs[name]; ok {
		templ = templ.Funcs(funcs)
	}

	return render.HTML{
		Template: templ,
		Data:     data,
	}
}

func (v *viewer) addNormalTempl(name string, paths ...string) {
	templs := make([]string, 0, len(paths)+1)

	templs = append(templs, "layouts/normal.html")
	templs = append(templs, paths...)

	v.templs[name] = templs
}

func (v *viewer) addMainTempl(name string, paths ...string) {
	templs := make([]string, 0, len(paths)+2)

	templs = append(templs, "layouts/main.html", "layouts/nav.html")
	templs = append(templs, paths...)

	v.templs[name] = templs
}

func (v *viewer) addTemplFunc(name string, funcs template.FuncMap) {
	v.funcs[name] = funcs
}

// NewRender return an render instance
func NewRender(box *rice.Box) render.HTMLRender {
	v := &viewer{
		box:    box,
		templs: make(HtmlTemplMap),
		funcs:  make(HtmlFuncMap),
	}

	v.addNormalTempl("error", "error.html")
	v.addNormalTempl("login", "login.html")
	v.addNormalTempl("password", "password.html")

	v.addMainTempl("home", "home.html")
	v.addMainTempl("user", "user/index.html", "user/search.html", "user/add.html", "user/edit.html")
	v.addMainTempl("role", "role/index.html")

	// TODO: v.addTemplFunc

	return v
}

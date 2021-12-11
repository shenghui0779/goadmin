package html

import (
	"html/template"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin/render"
)

type (
	HtmlTplMap  map[string][]string
	HtmlFuncMap map[string]template.FuncMap
)

type viewer struct {
	box   *rice.Box
	tpls  HtmlTplMap
	funcs HtmlFuncMap
}

// Instance supply render string
func (v *viewer) Instance(name string, data interface{}) render.Render {
	tpl := template.New(name)

	if htmls, ok := v.tpls[name]; ok {
		var builder strings.Builder

		for _, name := range htmls {
			builder.WriteString(v.box.MustString(name))
		}

		tpl = template.Must(tpl.Parse(builder.String()))
	}

	if funcs, ok := v.funcs[name]; ok {
		tpl = tpl.Funcs(funcs)
	}

	return render.HTML{
		Template: tpl,
		Data:     data,
	}
}

func (v *viewer) addNormalTemplate(name string, paths ...string) {
	tpls := make([]string, 0, len(paths)+1)

	tpls = append(tpls, "layouts/normal.html")
	tpls = append(tpls, paths...)

	v.tpls[name] = tpls
}

func (v *viewer) addMainTemplate(name string, paths ...string) {
	tpls := make([]string, 0, len(paths)+2)

	tpls = append(tpls, "layouts/main.html", "layouts/nav.html")
	tpls = append(tpls, paths...)

	v.tpls[name] = tpls
}

func (v *viewer) addTemplateFunc(name string, funcs template.FuncMap) {
	v.funcs[name] = funcs
}

// NewRender return an render instance
func NewRender(box *rice.Box) render.HTMLRender {
	v := &viewer{
		box:   box,
		tpls:  make(HtmlTplMap),
		funcs: make(HtmlFuncMap),
	}

	v.addNormalTemplate("error", "error.html")
	v.addNormalTemplate("login", "login.html")
	v.addNormalTemplate("password", "password.html")
	v.addMainTemplate("home", "home.html")
	v.addMainTemplate("user", "user/index.html", "user/search.html", "user/add.html", "user/edit.html")
	v.addMainTemplate("role", "role/index.html")

	// TODO: v.addTemplateFunc

	return v
}

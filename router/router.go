package router

import (
	// "elichika/webui"
	// "elichika/webui/user"

	"html/template"

	"github.com/gin-gonic/gin"
)

// Other packages should import the router package and declare which API it want to handle
// In main.go, import the relevant packages to support them
type HandlerInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
type SpecialGroupSetup = func(*gin.RouterGroup)
type GroupInfo struct {
	InitialHandlers []gin.HandlerFunc
	Handlers        map[string]HandlerInfo
	SpecialSetups   []SpecialGroupSetup
}

var (
	groups    = map[string]*GroupInfo{}
	templates []string
)

func initGroup(g string) {
	_, exist := groups[g]
	if exist {
		return
	}
	groups[g] = &GroupInfo{
		InitialHandlers: []gin.HandlerFunc{},
		Handlers:        map[string]HandlerInfo{},
		SpecialSetups:   []SpecialGroupSetup{},
	}
}

func (g *GroupInfo) AddInitialHandler(handler gin.HandlerFunc) {
	g.InitialHandlers = append(g.InitialHandlers, handler)
}

func (g *GroupInfo) AddHandler(method, path string, handler gin.HandlerFunc) {
	_, exist := g.Handlers[method+path]
	if exist {
		panic("Multiple handler for path and method: " + method + " " + path)
	}
	g.Handlers[method+path] = HandlerInfo{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

func (g *GroupInfo) AddSpecialSetup(specialGroupSetup SpecialGroupSetup) {
	g.SpecialSetups = append(g.SpecialSetups, specialGroupSetup)
}

func AddInitialHandler(group string, handler gin.HandlerFunc) {
	initGroup(group)
	groups[group].AddInitialHandler(handler)
}
func AddHandler(group, method, path string, handler gin.HandlerFunc) {
	initGroup(group)
	groups[group].AddHandler(method, path, handler)
}

func AddSpecialSetup(group string, specialGroupSetup SpecialGroupSetup) {
	initGroup(group)
	groups[group].AddSpecialSetup(specialGroupSetup)
}

func AddTemplates(path string) {
	templates = append(templates, path)
}

func Router(r *gin.Engine) {
	r.Static("/static", "static")
	r.StaticFile("/favicon.ico", "./webui/favicon.ico")

	funcs := template.FuncMap{}
	funcs["noescape"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	r.SetFuncMap(funcs)
	r.LoadHTMLFiles(templates...)

	for groupPath, groupInfo := range groups {
		groupApi := r.Group(groupPath, groupInfo.InitialHandlers...)
		for _, handlerInfo := range groupInfo.Handlers {
			switch handlerInfo.Method {
			case "POST":
				groupApi.POST(handlerInfo.Path, handlerInfo.Handler)
			case "GET":
				groupApi.GET(handlerInfo.Path, handlerInfo.Handler)
			default:
				panic("must be GET or POST only")
			}
		}
		for _, specialSetup := range groupInfo.SpecialSetups {
			specialSetup(groupApi)
		}
	}
}

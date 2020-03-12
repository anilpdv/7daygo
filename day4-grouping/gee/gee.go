package gee

import (
	"net/http"
)

// HandlerFunc : func (*Context)
// Context : check router.go
type HandlerFunc func(*Context)

// RouterGroup struct for structuring apis end points
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// Engine : struct
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New : creates a new engine instance
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// Group : func(prefix string)
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	// initialing group engine
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	// : appending group to the engine groups
	group.engine.groups = append(group.engine.groups, group)
	return newGroup
}

// GET : func(w http.ResponseWriter, r *http.Request)
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST : func(w http.ResponseWriter, r *http.Request)
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

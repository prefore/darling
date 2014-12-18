package darling

import (
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Params   []string
}

type Controller struct {
	Context *Context
	Data    map[interface{}]interface{}
	Name    string
}

type ControllerInterface interface {
	Init(c *Context, name string)
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
}

func (c *Controller) Init(context *Context, name string) {
	c.Context = context
	c.Name = name
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
	http.Error(c.Context.Response, "Method Not Allowed", 405)
}

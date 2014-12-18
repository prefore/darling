package darling

import (
	"net/http"
	"reflect"
	"regexp"
)

type ControllerInfo struct {
	controllerType reflect.Type
	regex          *regexp.Regexp
	handler        http.Handler
}

type ControllerRegistor struct {
	routers []*ControllerInfo
}

func NewControllerRegistor() *ControllerRegistor {
	return &ControllerRegistor{
		routers: make([]*ControllerInfo, 0),
	}
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {

	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		panic(regexErr)
		return
	}

	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()
	route := &ControllerInfo{}
	route.regex = regex
	route.controllerType = t
	p.routers = append(p.routers, route)
}

func (p *ControllerRegistor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var started bool
	requestPath := r.URL.Path

	for _, route := range p.routers {

		if !route.regex.MatchString(requestPath) {
			continue
		}

		matches := route.regex.FindStringSubmatch(requestPath)

		if len(matches[0]) != len(requestPath) {
			continue
		}

		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(&Context{Request: r, Response: rw, Params: matches[1:]})
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)
		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)
		if r.Method == "GET" {
			method = vc.MethodByName("Get")
			method.Call(in)
		} else if r.Method == "POST" {
			method = vc.MethodByName("Post")
			method.Call(in)
		} else if r.Method == "HEAD" {
			method = vc.MethodByName("Head")
			method.Call(in)
		} else if r.Method == "DELETE" {
			method = vc.MethodByName("Delete")
			method.Call(in)
		} else if r.Method == "PUT" {
			method = vc.MethodByName("Put")
			method.Call(in)
		} else if r.Method == "PATCH" {
			method = vc.MethodByName("Patch")
			method.Call(in)
		} else if r.Method == "OPTIONS" {
			method = vc.MethodByName("Options")
			method.Call(in)
		}
		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	if started == false {
		http.NotFound(rw, r)
	}
}

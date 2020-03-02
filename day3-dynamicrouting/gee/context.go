package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H type map[string]interface{}
type H map[string]interface{}

// Context struct
type Context struct {
	// origin begins
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param : get params in routes
// example /shoes/:id , we can id value
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm : gets body values
// example  : gets form values
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query : gets the query strings
// example : /something?q=state
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status : helps setting satus in the headers
// example: 404,500
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader : sets the header for the client
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String : helps sending values to the client
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON : json helps to send json data
func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data : sends data to the client
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML : data helps to send html code
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

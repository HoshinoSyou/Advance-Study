package webogo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Context struct {
	req        *http.Request
	res        http.ResponseWriter
	queryCache map[string]string
	formCache  map[string]string
	keys       map[string]interface{}
}

func InitContext(req *http.Request, res http.ResponseWriter) (c Context) {
	c = Context{
		req:        req,
		res:        res,
		queryCache: parseQuery(req.RequestURI),
		formCache:  make(map[string]string),
	}
	return
}

func parseQuery(uri string) (res map[string]string) {
	uris := strings.Split(uri, "?")
	if len(uris) == 1 {
		return
	}
	param := uris[len(uris)-1]
	params := strings.Split(param, "&")
	res = make(map[string]string)
	for _, v := range params {
		vs := strings.Split(v, "=")
		if len(vs) != 2 {
			fmt.Println(vs)
			panic("请求错误-uri错误")
		}
		res[vs[0]] = vs[1]
	}
	return
}

//func statusJudge(code int) bool {
//	switch {
//	case code >= 100 && code < 200:
//		return false
//	case code == http.StatusNoContent:
//		return false
//	case code == http.StatusNotModified:
//		return false
//	default:
//		return true
//	}
//}

//func (c *Context) status(code int) {
//	c.res.WriteHeader(code)
//	if statusJudge(code) {
//	}
//}

func (c *Context) String(str string) {
	_, _ = c.res.Write([]byte(str))
}

func (c *Context) JSON(code int, obj map[string]interface{}) {
	json, err := json.Marshal(&obj)
	if err != nil {
		log.Fatalf("序列化JSON格式失败喵！错误信息：%v", err)
	}
	c.res.WriteHeader(code)
	c.res.Write(json)
}

func (c *Context) XML(code int, obj map[string]interface{}) {
	xml, err := xml.Marshal(obj)
	if err != nil {
		log.Fatalf("序列化XML格式失败喵！错误信息：%v", err)
	}
	c.res.WriteHeader(code)
	c.res.Write(xml)
}

func (c *Context) Query(key string) string {
	val := c.queryCache[key]
	return val
}

func (c *Context) Get(key string) (val interface{}, bool bool) {
	val, bool = c.keys[key]
	return
}

func (c *Context) GetInt(key string) (i int) {
	val, ok := c.keys[key]
	if ok && val != nil {
		i = val.(int)
	}
	return
}

func (c *Context) GetString(key string) (s string) {
	val, ok := c.keys[key]
	if ok && val != nil {
		s = val.(string)
	}
	return s
}

func (c *Context) GetBool(key string) (b bool) {
	val, ok := c.keys[key]
	if ok && val != nil {
		b = val.(bool)
	}
	return b
}

func (c *Context) GetTime(key string) (t time.Time) {
	val, ok := c.keys[key]
	if ok && val != nil {
		t = val.(time.Time)
	}
	return t
}

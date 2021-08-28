package mux

import (
	"net/http"
)

type muxHandler struct{
	handlers map[string]http.Handler
	handlerFuncs map[string]func(resp http.ResponseWriter,req *http.Request)
}

func NewMuxHandler() *muxHandler{
	return &muxHandler{
		handlers: make(map[string]http.Handler),
		handlerFuncs: make(map[string]func(resp http.ResponseWriter,req *http.Request)),
	}
}

func (handler *muxHandler) ServeHTTP(resp http.ResponseWriter,req *http.Request){
	// 分发请求
	// 精确匹配
	urlPath := req.URL.Path
	if hl,ok := handler.handlers[urlPath];ok {
		hl.ServeHTTP(resp,req)
		return
	}

	if fn,ok := handler.handlerFuncs[urlPath];ok{
		fn(resp,req)
		return
	}
	// 找不到解析，执行NotFound
	http.NotFound(resp,req)
}

func (handler *muxHandler) Handler(pattern string,hl http.Handler){
	handler.handlers[pattern] = hl
}

func (handler *muxHandler) HandleFunc(pattern string,fn func(resp http.ResponseWriter,req *http.Request)){
	handler.handlerFuncs[pattern] = fn
}



package base

import (
	"net/http"
)

type Request struct {
	httpReq	*http.Request
	depth	uint32
}


func NewRequest(r *http.Request, depth uint32) *Request {
	return &Request{r, depth}
}

func (r *Request) HttpReq() *http.Request {
	return r.httpReq
}

func (r *Request) Depth() uint32 {
	return r.depth
}

type Response struct {
	httpResp *http.Response
	depth	uint32
}

func NewResponse(r *http.Response, depth uint32) *Response {
	return &Response{r, depth}
}



type Data interface {
	Valid()		bool
}

package apprequest

import "github.com/valyala/fasthttp"

var (
	POST            = []byte(fasthttp.MethodPost)
	GET             = []byte(fasthttp.MethodGet)
	PUT             = []byte(fasthttp.MethodPut)
	PATCH           = []byte(fasthttp.MethodPatch)
	DELETE          = []byte(fasthttp.MethodDelete)
	ApplicationJSON = []byte("application/json")
)

type HTTPRequest interface {
	NewRequest(body []byte, method []byte, url string) (*fasthttp.Request, *fasthttp.Response)
	FastSetHeaderAuthorizationBearer(req *fasthttp.Request, token string)
}
type fastHTTP struct {
}

func NewRequester() *fastHTTP { return &fastHTTP{} }

func (u *fastHTTP) NewRequest(body []byte, method []byte, url string) (*fasthttp.Request, *fasthttp.Response) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetBody(body)
	req.Header.SetMethodBytes(method)
	req.SetRequestURIBytes([]byte(url))
	return req, resp
}

func (u *fastHTTP) FastSetHeaderAuthorizationBearer(req *fasthttp.Request, token string) {
	req.Header.SetBytesKV([]byte("Authorization"), []byte("Bearer "+token))
}

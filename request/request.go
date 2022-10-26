package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RequestObj struct {
	u    string         // url
	p    url.Values     // params
	resp *http.Response // response
	err  error          // error
}

func New(u string) *RequestObj {
	return &RequestObj{
		u: u,
		p: url.Values{},
	}
}

func (r *RequestObj) SetParams(params map[string]interface{}) *RequestObj {
	for key, val := range params {
		switch val := val.(type) {
		case []string:
			for i, v := range val {
				r.p.Add(key+"["+strconv.Itoa(i)+"]", v)
			}
		default:
			r.p.Add(key, val.(string))
		}
	}
	return r
}

func (r *RequestObj) Get() *RequestObj {
	r.resp, r.err = http.Get(r.u + "?" + r.p.Encode())
	if r.err != nil {
		return r
	}
	return r
}

func (r *RequestObj) Post() *RequestObj {
	r.resp, r.err = http.PostForm(r.u, r.p)
	if r.err != nil {
		return r
	}
	return r
}

func (r *RequestObj) PostJSON() *RequestObj {
	paramsJson, err := json.Marshal(r.p)
	if err != nil {
		r.err = err
		return r
	}
	r.resp, r.err = http.Post(r.u, "application/json", bytes.NewBuffer(paramsJson))
	if r.err != nil {
		return r
	}
	return r
}

func (r *RequestObj) GetBody() ([]byte, error) {
	var res []byte
	if _, err := r.resp.Body.Read(res); err != nil {
		return nil, err
	}
	defer r.resp.Body.Close()
	return io.ReadAll(r.resp.Body)
}

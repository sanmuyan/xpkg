package xrequest

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Body []byte
	*http.Response
}

type Options struct {
	URL                string
	Method             string
	Body               []byte
	bodyReader         io.Reader
	Head               map[string]string
	Timeout            int
	InsecureSkipVerify bool
}

type Request struct {
	config *Options
}

func NewRequest(opt *Options) *Request {
	if opt == nil {
		opt = &Options{}
	}
	if opt.Timeout < 1 {
		opt.Timeout = 30
	}
	if opt.Body != nil {
		opt.bodyReader = bytes.NewReader(opt.Body)
	}
	return &Request{config: opt}
}

// Request 支持普通 HTTP 请求
func (c *Request) Request() (*Response, error) {
	req, err := http.NewRequest(c.config.Method, c.config.URL, c.config.bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range c.config.Head {
		req.Header.Set(k, v)
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: c.config.InsecureSkipVerify},
		},
		Timeout: time.Second * time.Duration(c.config.Timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Body: res, Response: resp}, nil
}

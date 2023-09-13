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
	Head               map[string]string
	Timeout            int
	InsecureSkipVerify bool
}

type Request struct {
	Config   *Options
	Response *Response
}

// Request 支持普通 HTTP 请求
func (c *Request) Request() (*Response, error) {
	req, err := http.NewRequest(c.Config.Method, c.Config.URL, bytes.NewReader(c.Config.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range c.Config.Head {
		req.Header.Set(k, v)
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: c.Config.InsecureSkipVerify},
		},
		Timeout: time.Second * time.Duration(c.Config.Timeout),
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

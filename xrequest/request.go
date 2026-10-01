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
	Client             *http.Client
	Request            *http.Request
	NoResponseBody     bool
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
	if opt.Client == nil {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 10000
		transport.MaxIdleConnsPerHost = 10000
		transport.MaxConnsPerHost = 10000
		transport.IdleConnTimeout = 60 * time.Second
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: opt.InsecureSkipVerify}
		opt.Client = &http.Client{
			Timeout:   time.Duration(opt.Timeout) * time.Second,
			Transport: transport,
		}
	}
	return &Request{config: opt}
}

// Request 支持普通 HTTP 请求
func (c *Request) Request() (*Response, error) {
	req := c.config.Request
	if req == nil {
		var err error
		req, err = http.NewRequest(c.config.Method, c.config.URL, c.config.bodyReader)
		if err != nil {
			return nil, err
		}
	}
	for k, v := range c.config.Head {
		req.Header.Set(k, v)
	}
	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.config.NoResponseBody {
		return &Response{Response: resp}, nil
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

func (c *Request) Client() *http.Client {
	return c.config.Client
}

package xrequest

import "testing"

func TestRequest(t *testing.T) {
	req := NewRequest(&Options{
		URL: "https://www.google.com",
	})
	res, err := req.Request()
	if err != nil {
		t.Error(err)
	}
	t.Logf("statusCode=%d", res.StatusCode)
}

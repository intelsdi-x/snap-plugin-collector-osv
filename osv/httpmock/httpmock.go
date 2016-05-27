// Simple HTTP mocking mechanism inspired by https://github.com/jarcoal/httpmock

package httpmock

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var Mock bool = false

type responder struct {
	reqType  string
	reqURL   string
	response string
	status   int
}

var responders = []responder{}

func RegisterResponder(reqType, reqURL, response string, status int) {
	newResponder := responder{reqType, reqURL, response, status}
	responders = append(responders, newResponder)
}

func ResetResponders() {
	responders = nil
}

func Get(url string) (resp *http.Response, err error) {
	if Mock {
		return createResponse(url, "GET")
	} else {
		return http.Get(url)
	}
}

func PostForm(url string, data url.Values) (resp *http.Response, err error) {
	if Mock {
		return createResponse(url, "POST")
	} else {
		return http.PostForm(url, data)
	}
}

func createResponse(url, reqType string) (resp *http.Response, err error) {
	for _, r := range responders {
		if r.reqURL == url && strings.ToUpper(r.reqType) == reqType {
			return &http.Response{
				StatusCode: r.status,
				Body:       responseBodyFromString(r.response),
			}, nil
		}
	}
	return nil, fmt.Errorf("URL %s not registered as responder for %s!", url, reqType)
}

func responseBodyFromString(body string) io.ReadCloser {
	return &dummyReadCloser{strings.NewReader(body)}
}

type dummyReadCloser struct {
	body io.ReadSeeker
}

func (d *dummyReadCloser) Read(p []byte) (n int, err error) {
	n, err = d.body.Read(p)
	if err == io.EOF {
		d.body.Seek(0, 0)
	}
	return n, err
}

func (d *dummyReadCloser) Close() error {
	return nil
}

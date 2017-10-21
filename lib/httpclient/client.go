package httpclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	Client = NewHTTPClient(http.DefaultClient)
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient(client *http.Client) *HTTPClient {
	if client == nil {
		client = http.DefaultClient
	}

	return &HTTPClient{
		client: client,
	}
}

func (hc *HTTPClient) SetHttpClient(client *http.Client) {
	hc.client = client
}

func (client *HTTPClient) UrlEncode(url string, params ...url.Values) string {
	if len(params) == 0 {
		return url
	}

	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}

	url += params[0].Encode()

	return url
}

func (hc *HTTPClient) Do(method, urlStr string, headers map[string]string, data ...interface{}) (*http.Response, error) {
	var (
		request *http.Request
		err     error
	)

	if len(data) == 0 {
		request, err = http.NewRequest(method, urlStr, nil)

	} else {
		var reader io.Reader

		body := data[0]
		switch body.(type) {
		case io.Reader:
			reader, _ = body.(io.Reader)

		case string:
			s, _ := body.(string)

			reader = bytes.NewBufferString(s)

		case []byte:
			buf, _ := body.([]byte)

			reader = bytes.NewBuffer(buf)

		case url.Values:
			params, _ := body.(url.Values)

			reader = bytes.NewBufferString(params.Encode())

		default:
			reader = strings.NewReader(fmt.Sprintf("%v", body))

		}

		request, err = http.NewRequest(method, urlStr, reader)
	}

	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := hc.client.Do(request)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode/100 != 2 {
		body, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			err = ioerr
		} else {
			resp.Body.Close()
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			err = errors.New(string(body))
		}
	}

	return resp, err
}

func (hc *HTTPClient) DoJSON(method string, url string, data interface{}, result interface{}) (*http.Response, error) {
	return hc.DoJSONWithHeaders(method, url, nil, data, result)
}

func (hc *HTTPClient) DoJSONWithHeaders(method string, url string, headers map[string]string, data interface{}, result interface{}) (*http.Response, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"

	resp, err := hc.Do(method, url, headers, b)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return resp, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if len(body) == 0 {
		return resp, nil
	}

	return resp, json.Unmarshal(body, &result)
}

func (hc *HTTPClient) DoXML(method string, url string, data interface{}, result interface{}) (*http.Response, error) {
	return hc.DoXMLWithHeaders(method, url, nil, data, result)
}

func (hc *HTTPClient) DoXMLWithHeaders(method, url string, headers map[string]string, data, result interface{}) (*http.Response, error) {
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "text/xml"

	resp, err := hc.Do(method, url, headers, b)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return resp, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if len(body) == 0 {
		return resp, nil
	}

	return resp, xml.Unmarshal(body, &result)
}

// NOTE: response body should be closed by callee
func (hc *HTTPClient) DoForm(method string, url string, data interface{}) (*http.Response, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	return hc.Do(method, url, headers, data)
}

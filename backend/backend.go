package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// RateLimit describes the current rate limit at which requests are being handled
type RateLimit struct {
	Limit     int64
	Remaining int64
	Reset     time.Duration
}

// Backend is used to Backendure the client backend
type Backend struct {
	HTTPClient *http.Client
	RateLimit  *RateLimit
	URL        string
	Version    string
}

type Response struct {
	OK        bool             `json:"OK"`
	Error     *statusCodeError `json:"error"`
	RateLimit *RateLimit       `json:"rateLimit"`
}

type nopReadCloser struct {
	io.Reader
}

func (nopReadCloser) Close() error { return nil }

func (r Response) Err() error {
	if r.OK {
		return nil
	}

	return errors.New(r.Error.Message)
}

// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code    int
	Status  string
	Message string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("%d error: %s. %s", t.Code, t.Status, t.Message)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

func FormatURLPath(format string, params ...string) string {
	// Convert parameters to interface{} and URL-escape them
	untypedParams := make([]interface{}, len(params))
	for i, param := range params {
		untypedParams[i] = interface{}(url.QueryEscape(param))
	}

	return fmt.Sprintf(format, untypedParams...)
}

func getRateLimit(r *http.Response) *RateLimit {
	limit, err := strconv.ParseInt(r.Header.Get("X-RateLimit-Limit"), 10, 64)
	if err != nil {
		return nil
	}
	remaining, err := strconv.ParseInt(r.Header.Get("X-RateLimit-Remaining"), 10, 64)
	if err != nil {
		return nil
	}
	reset, err := strconv.ParseInt(r.Header.Get("X-RateLimit-Reset"), 10, 64)
	if err != nil {
		return nil
	}
	return &RateLimit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     time.Duration(reset),
	}
}

func checkStatusCode(r *http.Response) error {
	if r.StatusCode == http.StatusTooManyRequests {
		return errors.New("Too many requests")
	}

	if r.StatusCode != http.StatusOK {
		return statusCodeError{Code: r.StatusCode, Status: r.Status}
	}

	return nil
}

// Call deserializes a requested url and binds it to an interface
func (b *Backend) Call(method, path, key string, v interface{}) error {
	req, err := b.NewRequest(method, path, key, "application/x-www-form-urlencoded")
	if err != nil {
		return err
	}

	if err := b.Do(req, nil, v); err != nil {
		return err
	}

	return nil
}

// CallBytes returns raw bytes from a call
func (b *Backend) CallBytes(method, path, key string) ([]byte, error) {
	req, err := b.NewRequest(method, path, key, "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	bytes, err := b.DoRaw(req, nil)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (b *Backend) NewRequest(method, path, key, contentType string) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = b.URL + path

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		fmt.Printf("Cannot create request: %v", err)
		return nil, err
	}

	authorization := "Bearer " + key

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("FloodIQ-Version", b.Version)

	return req, nil
}

func (b *Backend) Do(req *http.Request, body *bytes.Buffer, v interface{}) error {

	var res *http.Response
	var err error

	if body != nil {
		reader := bytes.NewReader(body.Bytes())

		req.Body = nopReadCloser{reader}

		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}

	res, err = b.HTTPClient.Do(req)

	if err != nil {
		fmt.Errorf("Request failed: %v", err)
		return err
	}

	defer res.Body.Close()

	b.RateLimit = getRateLimit(res)
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Errorf("Cannot read response: %v", err)
		return err
	}

	if res.StatusCode >= 400 {
		// SendError, res + resBody
		return nil
	}

	if v != nil {
		return parseResponseBody(resBody, v)
	}

	return nil
}

func parseResponseBody(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

func (b *Backend) DoRaw(req *http.Request, body *bytes.Buffer) ([]byte, error) {

	var res *http.Response
	var err error

	if body != nil {
		reader := bytes.NewReader(body.Bytes())

		req.Body = nopReadCloser{reader}

		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}

	res, err = b.HTTPClient.Do(req)

	if err != nil {
		fmt.Errorf("Request failed: %v", err)
		return nil, err
	}

	defer res.Body.Close()

	b.RateLimit = getRateLimit(res)
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Errorf("Cannot read response: %v", err)
		return nil, err
	}

	if res.StatusCode >= 400 {
		fmt.Errorf("Cannot find resource: %v", err)
		// SendError, res + resBody
		return nil, err
	}

	return resBody, nil
}

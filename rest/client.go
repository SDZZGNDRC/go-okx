package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/SDZZGNDRC/go-okx/rest/api"

	"github.com/SDZZGNDRC/go-okx/common"

	"github.com/google/go-querystring/query"
)

var (
	specialPath = [...]string{ // these api require signature
		"/api/v5/account/balance",
		"/api/v5/account/config",
		"/api/v5/account/positions",
		"/api/v5/finance/staking-defi/orders-active",
		"/api/v5/finance/staking-defi/offers",
		"/api/v5/trade/orders-pending",
		"/api/v5/trade/order",
		"/api/v5/trade/order",
	}
)

type Client struct {
	Host string
	Auth common.Auth
	C    *http.Client
}

// new *Client
func New(host string, auth common.Auth, c *http.Client) *Client {
	if host == "" {
		host = "https://www.okx.com"
	}
	if c == nil {
		c = http.DefaultClient
	}

	return &Client{
		Host: host,
		Auth: auth,
		C:    c,
	}
}

// do request
// Use package net/http
func (c *Client) Do2(req api.IRequest, resp api.IResponse) error {
	data, err := c.do2(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	if !resp.IsOk() {
		return fmt.Errorf("code: %s, message: %s", resp.GetCode(), resp.GetMessage())
	}

	return nil
}

// for historical dealprices data
func (c *Client) Do3(req api.IRequest, FilePath string) error {
	// data, err := c.do3(req)
	err := c.DumpFile(req, FilePath)
	return err
}

// do2 request
// Use package net/http
func (c *Client) do2(r api.IRequest) ([]byte, error) {
	var body []byte
	var err error
	req := c.newRequest2(r)
	resp, err := c.C.Do(req)
	if err != nil { // ======================
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		return nil, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode, string(body))
	}
	body, _ = ioutil.ReadAll(resp.Body)
	return body, nil
}

func (c *Client) DumpFile(r api.IRequest, FilePath string) error {
	var body []byte
	var err error
	req := c.newRequest3(r)
	resp, err := c.C.Do(req)
	if err != nil { // ======================
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode, string(body))
	}
	file, err := os.Create(FilePath)
	if err != nil {
		fmt.Println("[Error] Can't create file: ", err.Error())
		os.Exit(-1)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// new Request
// for package net/http
func (c *Client) newRequest2(r api.IRequest) *http.Request {
	sign := c.newSignature(r)
	var body io.Reader = nil
	headers := map[string]string{
		"Content-Type":         "application/json;charset=utf-8",
		"Accept":               "application/json",
		"OK-ACCESS-KEY":        c.Auth.ApiKey,
		"OK-ACCESS-PASSPHRASE": c.Auth.Passphrase,
		"OK-ACCESS-SIGN":       sign.Build(),
		"OK-ACCESS-TIMESTAMP":  sign.Timestamp,
	}
	if c.Auth.Simulated {
		headers["x-simulated-trading"] = "1"
	}
	if sign.Body != "" {
		body = strings.NewReader(sign.Body)
		// body = ioutil.NopCloser(strings.NewReader(sign.Body))
	}
	req, err := http.NewRequest(r.GetMethod(), c.Host+sign.Path, body)
	if err != nil {
		panic(err)
	}
	if isSpecialPath(r.GetPath()) {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		// fmt.Println("[DEBUG] Add Signature")
	}

	// to solve to EOF problem
	// Refer to https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi
	req.Close = true

	return req
}

func (c *Client) newRequest3(r api.IRequest) *http.Request {
	sign := c.newSignature(r)
	var body io.Reader = nil
	headers := map[string]string{
		"Content-Type":         "application/json;charset=utf-8",
		"Accept":               "application/json",
		"OK-ACCESS-KEY":        c.Auth.ApiKey,
		"OK-ACCESS-PASSPHRASE": c.Auth.Passphrase,
		"OK-ACCESS-SIGN":       sign.Build(),
		"OK-ACCESS-TIMESTAMP":  sign.Timestamp,
	}
	if c.Auth.Simulated {
		headers["x-simulated-trading"] = "1"
	}
	if sign.Body != "" {
		body = strings.NewReader(sign.Body)
		// body = ioutil.NopCloser(strings.NewReader(sign.Body))
	}
	req, err := http.NewRequest(r.GetMethod(), "https://static.okx.com"+sign.Path, body)
	if err != nil {
		panic(err)
	}
	if isSpecialPath(r.GetPath()) {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		// fmt.Println("[DEBUG] Add Signature")
	}

	// to solve to EOF problem
	// Refer to https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi
	req.Close = true

	return req
}

// new *Signature
func (c *Client) newSignature(r api.IRequest) *common.Signature {
	var body []byte
	path := r.GetPath()

	if r.IsPost() {
		body, _ = json.Marshal(r.GetParam())
	} else if values, _ := query.Values(r.GetParam()); len(values) > 0 {
		path += "?" + values.Encode()
	}

	return c.Auth.Signature(r.GetMethod(), path, string(body), false)
}

// Special Path should contain signature.
func isSpecialPath(path string) bool {
	for _, b := range specialPath {
		if b == path {
			return true
		}
	}
	return false
}

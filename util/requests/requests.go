package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	urlUtil "net/url"
	"time"
)

// GET, send get req
func GET(url string, params map[string]string, headers map[string]string, timeout int) (response []byte) {
	if params != nil {
		var rq = urlUtil.Values{}
		for k, v := range params {
			rq.Add(k, v)
		}
		url = url + "?" + rq.Encode()
	}

	return sendRequest("GET", url, params, nil, headers, timeout)
}

// POST, send post req
func POST(url string, data interface{}, headers map[string]string, timeout int) (response []byte) {
	return sendRequest("POST", url, nil, data, headers, timeout)
}

// sendRequest, send request and receive response
func sendRequest(method, url string, params map[string]string, data interface{}, headers map[string]string, timeout int) (response []byte) {

	// create req
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	defer req.Body.Close()

	// send req
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// receive resp
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.Bytes()
}

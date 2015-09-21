package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
import "C"

type Command struct {
	URIs []string `json:"uris"`
}

type Result map[string]*Response
type Response struct {
	Body   string `json:"body"`
	Status int    `json:"status"`
	Err    string `json:"error"`
	Uri    string `json:"uri"`
}

func makeRequest(uri string) *Response {
	resp, err := http.Get(uri)
	if err != nil {
		return &Response{Status: resp.StatusCode}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{Status: resp.StatusCode, Err: err.Error()}
	}

	return &Response{Uri: uri, Body: string(body), Status: resp.StatusCode}
}

//export scatter_request
func scatter_request(data *C.char) string {
	cmd := Command{}
	err := json.Unmarshal([]byte(C.GoString(data)), &cmd)
	if err != nil {
		return err.Error()
	}

	c := make(chan *Response)
	for _, uri := range cmd.URIs {
		go func(_uri string) {
			c <- makeRequest(_uri)
		}(uri)
	}

	result := Result{}
	for i := 0; i < len(cmd.URIs); i++ {
		resp := <-c
		result[resp.Uri] = resp
	}

	close(c)

	b, err := json.Marshal(result)
	if err != nil {
		return "{}"
	}

	return string(b)
}

func main() {}
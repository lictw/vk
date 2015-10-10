package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const api = `https://api.vk.com/method/`

type Api struct {
	Token string
}

// "method" like "messages.get"
// All available method see here - https://vk.com/dev/methods
// "params" - parameters of the selected API method
func (this Api) Request(method string, params map[string]string) (result json.RawMessage, e error) {

	request := api + method + "?"
	for key, value := range params {
		request += key + "=" + value + "&"
	}
	request += "access_token=" + this.Token

	response, e := http.Get(request)
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()

	result, e = ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, e
	}

	var j js
	json.Unmarshal(result, &j)
	if j.Error.Code != 0 {
		return nil, errors.New(fmt.Sprint("vk: ", j.Error.Code, ", \"", j.Error.Message, "\""))
	}

	return j.Response, nil
}

type js struct {
	Error    vkError         `json:"error"`
	Response json.RawMessage `json:"response"`
}
type vkError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

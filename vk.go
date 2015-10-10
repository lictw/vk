package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const api = `https://api.vk.com/method/`

type js struct {
	Error    vkError         `json:"error"`
	Response json.RawMessage `json:"response"`
}

type vkError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

type Api struct {
	Token string
}

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
	if e := json.Unmarshal(result, &j); e != nil {
		return nil, e
	}

	if j.Error.Code != 0 {
		return nil, errors.New(fmt.Sprint("vk: ", j.Error.Code, ", \"", j.Error.Message, "\""))
	}

	return j.Response, nil
}

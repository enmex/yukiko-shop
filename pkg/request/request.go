package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Code int
	Body *string
}

func Post(url string, request interface{}) (*Response, error) {
	requestJson, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(requestJson),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := string(jsonData)

	return &Response{
		Code: res.StatusCode,
		Body: &jsonBody,
	}, nil
}

func Get(url string) (*Response, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := string(jsonData)

	return &Response{
		Code: res.StatusCode,
		Body: &jsonBody,
	}, nil
}

func Delete(url string) (*Response, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := string(jsonData)

	return &Response{
		Code: res.StatusCode,
		Body: &jsonBody,
	}, nil
}

func Patch(url string, request interface{}) (*Response, error) {
	requestJson, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(requestJson))

	if err != nil {
		return nil, nil
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, nil
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := string(jsonData)

	return &Response{
		Code: res.StatusCode,
		Body: &jsonBody,
	}, nil
}

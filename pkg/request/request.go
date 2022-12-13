package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Post(url string, request interface{}) (string, error) {
	requestJson, err := json.Marshal(request)

	if err != nil {
		return "", err
	}

	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(requestJson),
	)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	jsonBody := string(jsonData)

	return jsonBody, nil
}

func Get(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	jsonBody := string(jsonData)

	return jsonBody, nil
}

func Delete(url string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	jsonBody := string(jsonData)

	return jsonBody, nil
}

func Patch(url string, request interface{}) (string, error) {
	requestJson, err := json.Marshal(request)

	if err != nil {
		return "", err
	}

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(requestJson))

	if err != nil {
		return "", nil
	}

	res, err := client.Do(req)

	if err != nil {
		return "", nil
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	jsonBody := string(jsonData)

	return jsonBody, nil
}

package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"svna/src/constants"
)

func Get(URL string) ([]byte, error) {

	reqUrl, err := url.Parse(URL)

	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{
			"Authorization": {"Bearer " + constants.ACCESS_TOKEN},
		},
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return data, nil
}

func Post(postUrl string, jsonString string) ([]byte, error) {
	Print("Sending post to  " + postUrl)
	Print("Json body: " + jsonString)

	reqBody := strings.NewReader(jsonString)

	res, err := http.Post(postUrl, "application/json; charset=UTF-8", reqBody)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return data, nil
}

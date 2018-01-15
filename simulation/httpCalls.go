package simulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(url, token string) (body []byte, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming new request GET: %s", err)
	}

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil,
			fmt.Errorf("Status code of response indicates request failure(url: %s.  Code: %s",
				url,
				res.Status,
			)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return
}

func post(url, token string, data interface{}) (body []byte, err error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling json data: %s",
			err,
		)
	}

	// fmt.Println(jsonData)

	payload := strings.NewReader(string(jsonData))

	// fmt.Println(payload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code %d", res.StatusCode)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Cannot print out body with reasons for call failure: ", err)
	}

	return body, nil
}

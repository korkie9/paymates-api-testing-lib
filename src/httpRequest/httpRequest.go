package http

import (
	"bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"io"
	"net/http"
	"paymates-mock-db-updater/src/check_error"
	"paymates-mock-db-updater/src/util/env"
)

func HttpRequest(requesttype string, requestbody map[string]interface{}, path string, authorizationheader string) ([]byte, error, string) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	postBody, err := json.Marshal(requestbody)
	check_error.ErrCheck(err)

	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(requesttype, util.DotEnvVariable("API_URL")+path, requestBody)

	fmt.Println(util.DotEnvVariable("API_URL") + path)
	check_error.ErrCheck(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authorizationheader)

	resp, err := client.Do(req)
	check_error.ErrCheck(err)
	fmt.Println("Status: ", resp.StatusCode)

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	check_error.ErrCheck(err)

	return responseBody, nil, resp.Status
}

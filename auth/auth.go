package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paymates-mock-db-updater/checkError"
	"paymates-mock-db-updater/env"
)

type AuthUser struct {
	Uid                string `json:"uid"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	PhotoUrl           string `json:"photoUrl"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenExpiry string `json:"refreshTokenExpiry"`
}

func LoginAndGetToken() string {
	return "hello acess token"
}

func GetAccessToken() string {
	var user AuthUser = loginAndGetUser()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	postBody, _ := json.Marshal(map[string]string{
		"uid":                user.Uid,
		"refreshToken":       user.RefreshToken,
		"refreshTokenExpiry": user.RefreshTokenExpiry,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", env.DotEnvVariable("API_URL")+"Auth/refresh-token", responseBody)
	checkError.ErrCheck(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	checkError.ErrCheck(err)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	checkError.ErrCheck(err)
	token := string(body)
	fmt.Println("token: ", token)
	return token
}

func GetRefreshToken() string {
	fmt.Println("refresh token eskettiiiiiiiiit")
	return "hello access token"
}

func loginAndGetUser() AuthUser {
	//put this is another function
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	postBody, _ := json.Marshal(map[string]string{
		"username": "string",
		"password": "string",
	})
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", env.DotEnvVariable("API_URL")+"Auth/login", responseBody)
	checkError.ErrCheck(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	checkError.ErrCheck(err)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	checkError.ErrCheck(err)
	var user AuthUser
	err = json.Unmarshal(body, &user)
	checkError.ErrCheck(err)
	sb := string(body)
	fmt.Println(sb)
	return user
}

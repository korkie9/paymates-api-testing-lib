package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
	util "paymates-mock-db-updater/src/util/get_input"
	"regexp"
	"strconv"
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

type BaseResponse[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

// TODO: Adjust for base url
func GetAccessToken() string {
	user := LoginAndGetMockUser()

	requestBody := map[string]interface{}{
		"uid":                user.Data.Uid,
		"refreshToken":       user.Data.RefreshToken,
		"refreshTokenExpiry": user.Data.RefreshTokenExpiry,
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/access-token", "")
	check_error.ErrCheck(err)
	var tokenRes BaseResponse[string]
	_ = json.Unmarshal(res, &tokenRes)
	token := tokenRes.Data
	fmt.Println("Acess token: ", token)
	return token
}

func GetRefreshToken() string {
	fmt.Println("refresh token eskettiiiiiiiiit")
	return "hello access token"
}

func LoginAndGetMockUser() BaseResponse[AuthUser] {
	fmt.Println("Enter user name")
	username, err := util.GetUserInput()
	check_error.ErrCheck(err)
	fmt.Println("Enter user password")
	userPasssword, err := util.GetUserInput()
	check_error.ErrCheck(err)
	var user BaseResponse[AuthUser]
	requestBody := map[string]interface{}{
		"username": username,
		"password": userPasssword,
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/login", "")
	check_error.ErrCheck(err)
	_ = json.Unmarshal(res, &user)
	fmt.Println("refresh token: ", user.Data.RefreshToken)
	return user
}

func TestAPI() {
	var accessToken string = GetAccessToken()
	requestBody := map[string]interface{}{}
	res, err, _ := reqres.HttpRequest("GET", requestBody, "Auth/test/", accessToken)
	check_error.ErrCheck(err)
	testString := string(res)
	fmt.Println("Test response: ", testString)
}

func RegisterUser() AuthUser {
	fmt.Println("Enter user name")
	username, err := util.GetUserInput()
	check_error.ErrCheck(err)
	fmt.Println("Enter user email")
	userEmail, err := util.GetUserInput()
	check_error.ErrCheck(err)
	var user AuthUser
	requestBody := map[string]interface{}{
		"email":     userEmail,
		"photoUrl":  "string",
		"username":  username,
		"firstName": "string",
		"lastName":  "string",
		"password":  "string",
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/register", "")
	check_error.ErrCheck(err)
	resString := string(res)
	fmt.Println("Response body: ", resString)
	_ = json.Unmarshal(res, &user)
	fmt.Println("Created user with username: ", user.Username)
	fmt.Println("Check your mail")
	return user
}

func CreateUser() AuthUser {
	fmt.Println("Enter token")
	token, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"token": token,
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/create-user", "")
	check_error.ErrCheck(err)
	var user AuthUser
	_ = json.Unmarshal(res, &user)
	fmt.Println("Created user with username: ", user.Username)
	return user
}

func RegisterMultipleMockUsers(db *sql.DB) {
	fmt.Println("How many mock users would you like to add")
	// Read input
	var amount int
	for {
		input, err := util.GetUserInput()
		check_error.ErrCheck(err)
		// check if input matches an int value
		numberreg := regexp.MustCompile(`\d`)
		if !numberreg.MatchString(input) {
			fmt.Println("Please add a valid number")
		} else {
			castedAmount, err := strconv.Atoi(input)
			check_error.ErrCheck(err)
			amount = castedAmount
			break
		}

	}
	// register mock user
	for i := 0; i < amount; i++ {
		RegisterUser()
	}
}

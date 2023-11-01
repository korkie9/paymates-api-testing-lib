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

func GetAccessToken() string {
	var user AuthUser = LoginAndGetMockUser()

	requestBody := map[string]interface{}{
		"uid":                user.Uid,
		"refreshToken":       user.RefreshToken,
		"refreshTokenExpiry": user.RefreshTokenExpiry,
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/refresh-token", "")

	check_error.ErrCheck(err)
	accessToken := string(res)
	fmt.Println("Acess token: ", accessToken)
	return accessToken
}

func GetRefreshToken() string {
	fmt.Println("refresh token eskettiiiiiiiiit")
	return "hello access token"
}

func LoginAndGetMockUser() AuthUser {
	var user AuthUser
	requestBody := map[string]interface{}{
		"username": "string0",
		"password": "string",
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/login", "")
	check_error.ErrCheck(err)
	err = json.Unmarshal(res, &user)
	fmt.Println("refresh token: ", user.RefreshToken)
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

func RegisterMockUser(db *sql.DB) AuthUser {
	var count int
	_ = db.QueryRow(`select count(*) from Users`).Scan(&count)

	fmt.Println(count)
	mockEmail := "test" + strconv.Itoa(count) + "@test.com"
	mockUsername := "string" + strconv.Itoa(count)
	var user AuthUser
	requestBody := map[string]interface{}{
		"email":     mockEmail,
		"photoUrl":  "string",
		"username":  mockUsername,
		"firstName": "string",
		"lastName":  "string",
		"password":  "string",
	}
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Auth/register", "")
	check_error.ErrCheck(err)
	resString := string(res)
	fmt.Println("Response body: ", resString)
	err = json.Unmarshal(res, &user)
	fmt.Println("Created user with username: ", user.Username)
	return user
}

func RegisterMultipleMockUsers(db *sql.DB) {
	fmt.Println("How many mock users would you like to add")
	//Read input
	var amount int
	for {
		input, err := util.GetUserInput()
		check_error.ErrCheck(err)
		//check if input matches an int value
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
	//register mock user
	for i := 0; i < amount; i++ {
		RegisterMockUser(db)
	}
}

package users

import (
	"database/sql"
	"fmt"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
	util "paymates-mock-db-updater/src/util/get_input"
	truncate "paymates-mock-db-updater/src/util/truncate"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid                string `json:"uid"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Email              string `json:"email"`
	PhotoUrl           string `json:"photoUrl"`
	Username           string `json:"username"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenExpiry string `json:"refreshTokenExpiry"`
	Password           string `json:"Password"`
	Verified           string `json:"Verified"`
}

var usersList = []string{"Zac", "Amy", "Luke", "Justin", "Migs", "Micah", "Jade", "Aiden", "Frans"}

func TruncateUsers(db *sql.DB) {
	_, err := truncate.Truncate(db, "Users")
	check_error.ErrCheck(err)
}

func GetAllUsers(db *sql.DB) {
	fmt.Println("Getting users")
	users, err := db.Query("SELECT * FROM paymates.Users")
	check_error.ErrCheck(err)
	println("USERS:")
	for users.Next() {
		var user User
		err = users.Scan(&user.Uid, &user.FirstName, &user.LastName, &user.Email, &user.PhotoUrl, &user.Username, &user.RefreshToken, &user.RefreshTokenExpiry, &user.Password, &user.Verified)
		check_error.ErrCheck(err)
		fmt.Println(user) //, " ", user.FirstName, " ", user.LastName, " ", user.Email, " ", user.Username, " ", user.Password, " ", user.RefreshToken)
	}
}

func CreateUserMocks(db *sql.DB) {
	for index, user := range usersList {
		indexStr := fmt.Sprint(index)
		mockVal := "testing" + indexStr
		_, err := db.Exec(`Insert Into Users ( Uid, FirstName, LastName, Email, PhotoUrl, Username)
        values ( ?, ?, ?, ?, ?, ?)`, mockVal, user, user, mockVal+"@test.com", mockVal, mockVal)
		check_error.ErrCheck(err)
	}
}

func GetNumberOfUsers(db *sql.DB) int {
	fmt.Println("Getting amount users")
	var err error
	query := "SELECT COUNT(*) FROM Users"
	var rowCount int
	err = db.QueryRow(query).Scan(&rowCount)
	check_error.ErrCheck(err)
	fmt.Printf("Number of rows in the Users table: %d\n", rowCount)
	return rowCount
}

func GetUser() {
	fmt.Println("Enter the uid of the user you'd like to Find...")
	userId, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"uid": userId,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "User/get-user", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Get Friend Response: ", resStr)
}

func Test() {
	requestBody := map[string]interface{}{}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "User/get-user", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Test Friend Response: ", resStr)
}

package users

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"paymates-mock-db-updater/src/check_error"
	truncate "paymates-mock-db-updater/src/util/truncate"
)

type User struct {
	Uid       string `json:"uid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	PhotoUrl  string `json:"photoUrl"`
	Username  string `json:"username"`
}

var usersList = []string{"Zac", "Amy", "Luke", "Justin", "Migs", "Micah", "Jade", "Aiden", "Frans"}

func TruncateUsers(db *sql.DB) {
	_, err := truncate.Truncate(db, "Users")
	check_error.ErrCheck(err)
}

func GetAllUsers(db *sql.DB) {
	fmt.Println("Getting users")
	users, err := db.Query("SELECT * FROM Users")
	check_error.ErrCheck(err)
	println("USERS:")
	for users.Next() {
		var user User
		err = users.Scan(&user.Uid, &user.FirstName, &user.LastName, &user.Email, &user.PhotoUrl, &user.Username)
		check_error.ErrCheck(err)
		fmt.Println(user.Uid, " ", user.FirstName, " ", user.LastName, " ", user.Email, " ", user.Username)
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

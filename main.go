package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	"paymates-mock-db-updater/src/friends"
	"paymates-mock-db-updater/src/users"
	"paymates-mock-db-updater/src/util/env"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type User struct {
	Uid                string `json:"uid"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Email              string `json:"email"`
	PhotoUrl           string `json:"photoUrl"`
	Username           string `json:"username"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenExpiry string `json:"refreshTokenExpiry"`
}

func main() {
	connectionstring := util.DotEnvVariable("CONNECTION_STRING")
	//open sql connection
	db, err = sql.Open("mysql", connectionstring)
	check_error.ErrCheck(err)
	err = db.Ping()
	check_error.ErrCheck(err)
	fmt.Println("Connected to db")
	//Get user input command
	reader := bufio.NewReader(os.Stdin)
	printAvailibleCommands()
	for {
		//perforn function based on user input command
		input, err := reader.ReadString('\n')
		check_error.ErrCheck(err)
		input = strings.TrimSpace(input)
		switch input {
		case "clear users":
			users.TruncateUsers(db)
		case "get users":
			users.GetAllUsers(db)
		case "create mock users":
			users.GetAllUsers(db)
		case "register mock user":
			auth.RegisterMockUser(db)
		case "register multiple mock users":
			auth.RegisterMultipleMockUsers(db)
		case "get refresh token":
			auth.GetRefreshToken()
		case "get access token":
			var accessToken = auth.GetAccessToken()
			fmt.Println("access token: ", accessToken)
		case "get friends":
			friends.GetAllFriends(db)
		case "add friend":
			friends.AddFriend(db, "b4cce400-afd3-4a83-8f47-184d456cacb9")
		case "create mock friends":
			friends.CreateFriendsMocks(db)
		case "test auth":
			auth.TestAPI()
		default:
			fmt.Println("OOPS!! ", input, " is not an avaiible command")
			printAvailibleCommands()
		}
	}
}

func printAvailibleCommands() {
	fmt.Println("AVAILIBLE COMMANDS:")
	fmt.Println("==========================================")
	fmt.Println("get users; clear users; create mock users; register multiple mock users")
	fmt.Println("get refresh token; get access token;")
	fmt.Println("get friends; create mock friends; add friend;")
	fmt.Println("==========================================")
}

//Later for getting friend data
// SELECT U2.*
// FROM Users AS U1
// INNER JOIN Friends AS F ON U1.Uid = F.FriendOneUid
// INNER JOIN Users AS U2 ON U2.Uid = F.FriendTwoUid
// WHERE U1.Uid = 'testing1';

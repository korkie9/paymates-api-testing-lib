package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"paymates-mock-db-updater/auth"
	"paymates-mock-db-updater/checkError"
	"paymates-mock-db-updater/env"
	"paymates-mock-db-updater/friends"
	"paymates-mock-db-updater/users"
	"strings"
)

var db *sql.DB
var err error

func main() {
	connectionstring := env.DotEnvVariable("CONNECTION_STRING")
	//open sql connection
	db, err = sql.Open("mysql", connectionstring)
	checkError.ErrCheck(err)
	err = db.Ping()
	checkError.ErrCheck(err)
	fmt.Println("Connected to db")
	//Get user input command
	reader := bufio.NewReader(os.Stdin)
	printAvailibleCommands()
	for {
		//perforn function based on user input command
		input, err := reader.ReadString('\n')
		checkError.ErrCheck(err)
		input = strings.TrimSpace(input)
		switch input {
		case "clear users":
			users.TruncateUsers(db)
		case "get users":
			users.GetAllUsers(db)
		case "create mock users":
			users.GetAllUsers(db)
		case "get refresh token":
			auth.GetRefreshToken()
		case "get access token":
			auth.GetAccessToken()
		case "get friends":
			friends.GetAllFriends(db)
		case "create mock friends":
			friends.CreateFriendsMocks(db)
		default:
			fmt.Println("OOPS!! ", input, " is not an avaiible command")
			printAvailibleCommands()
		}
	}
	// defer db.Close()
}

func printAvailibleCommands() {
	fmt.Println("AVAILIBLE COMMANDS:")
	fmt.Println("==========================================")
	fmt.Println("get users; clear users; create mock users;")
	fmt.Println("get refresh token; get access token;")
	fmt.Println("get friends; create mock friends")
	fmt.Println("==========================================")
}

//Later for getting friend data
// SELECT U2.*
// FROM Users AS U1
// INNER JOIN Friends AS F ON U1.Uid = F.FriendOneUid
// INNER JOIN Users AS U2 ON U2.Uid = F.FriendTwoUid
// WHERE U1.Uid = 'testing1';

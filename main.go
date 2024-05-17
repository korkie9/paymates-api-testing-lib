package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	accounts "paymates-mock-db-updater/src/accounts"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	"paymates-mock-db-updater/src/friends"
	"paymates-mock-db-updater/src/transactions"
	"paymates-mock-db-updater/src/users"
	util "paymates-mock-db-updater/src/util/env"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	connectionstring := util.DotEnvVariable("CONNECTION_STRING")
	// open sql connection
	db, err = sql.Open("mysql", connectionstring)
	check_error.ErrCheck(err)
	err = db.Ping()
	check_error.ErrCheck(err)
	fmt.Println("Connected to db")
	// Get user input command
	reader := bufio.NewReader(os.Stdin)
	printAvailibleCommands()
	for {
		// perforn function based on user input command
		input, err := reader.ReadString('\n')
		check_error.ErrCheck(err)
		input = strings.TrimSpace(input)
		switch input {
		case "truncate users":
			users.TruncateUsers(db)
		case "get users":
			users.GetAllUsers(db)
		case "get number of users":
			users.GetNumberOfUsers(db)
		case "create mock users":
			users.GetAllUsers(db)
		case "register user":
			auth.RegisterUser()
		case "update user":
			users.UpdateUser()
		case "register multiple mock users":
			auth.RegisterMultipleMockUsers(db)
		case "get refresh token":
			auth.GetRefreshToken()
		case "get access token":
			accessToken := auth.GetAccessToken()
			fmt.Println("access token: ", accessToken)
		case "get all friends in db":
			friends.GetAllFriends(db)
		case "add friend":
			friends.AddFriend()
		case "add multiple friends":
			friends.AddMultipleFriends()
		case "create mock friends":
			friends.CreateFriendsMocks(db)
		case "truncate friends":
			friends.TruncateFriends(db)
		case "delete friend":
			friends.DeleteFriend(db)
		case "get user friends":
			friends.GetUserFriends()
		case "find friend":
			friends.FindFriend()
		case "test auth":
			auth.TestAPI()
		case "help":
			printAvailibleCommands()
		case "get user":
			users.GetUser()
		case "create user":
			auth.CreateUser()
		case "create transaction":
			transactions.CreateTransAction()
		case "login user":
			auth.LoginAndGetMockUser()
		case "get all transactions from db":
			transactions.GetAllTransactions(db)
		case "delete transaction":
			transactions.DeleteTransAction()
		case "print slice":
			transactions.PrintArray()
		case "remove account":
			accounts.DeleteAccount()
		case "upload photo":
			users.UploadPhoto()
		default:
			fmt.Println("OOPS!! ", input, " is not an avaiible command")
			printAvailibleCommands()
		}
	}
}

func printAvailibleCommands() {
	fmt.Println("AVAILIBLE COMMANDS:")
	fmt.Println("==========================================")
	fmt.Println("upload photo; get users; truncate users; create mock users; register multiple mock users; register user; get number of users; login user; update user;")
	fmt.Println("get refresh token; get access token; create user")
	fmt.Println("create transaction; get all transactions from db; delete transaction")
	fmt.Println("get all friends in db; create mock friends; add friend; truncate friends; delete friend; get user friends; find friend")
	fmt.Println("remove account")
	fmt.Println("==========================================")
}

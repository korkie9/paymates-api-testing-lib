package friends

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
)

type Friend struct {
	FriendOneUid string `json:"friendOneUid"`
	FriendTwoUid string `json:"friendTwoUid"`
}

var usersList = []string{"Zac", "Amy", "Luke", "Justin", "Migs", "Micah", "Jade", "Aiden", "Frans"}

func GetAllFriends(db *sql.DB) {
	users, err := db.Query("SELECT * FROM Friends")
	check_error.ErrCheck(err)
	println("FRIENDS:")
	for users.Next() {
		var friend Friend
		err = users.Scan(&friend.FriendOneUid, &friend.FriendTwoUid)
		check_error.ErrCheck(err)

		fmt.Println(friend.FriendOneUid, " ", friend.FriendTwoUid)
	}
}

func CreateFriendsMocks(db *sql.DB) {
	friendOne := "testing2"
	for index := range usersList {
		indexStr := fmt.Sprint(index)
		mockVal := "testing" + indexStr
		if mockVal != friendOne {
			_, err := db.Exec(`Insert Into Friends ( FriendOneUid, FriendTwoUid)
	        values ( ?, ?)`, friendOne, mockVal)
			check_error.ErrCheck(err)

		}
	}
}

func AddFriend(db *sql.DB, friendUid string) {
	requestBody := map[string]string{
		"friendUid": friendUid,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Friends/add-friend", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Friend Response: ", resStr)
}

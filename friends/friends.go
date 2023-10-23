package friends

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"paymates-mock-db-updater/checkError"
)

type Friend struct {
	FriendOneUid string `json:"friendOneUid"`
	FriendTwoUid string `json:"friendTwoUid"`
}

var usersList = []string{"Zac", "Amy", "Luke", "Justin", "Migs", "Micah", "Jade", "Aiden", "Frans"}

func GetAllFriends(db *sql.DB) {
	users, err := db.Query("SELECT * FROM Friends")
	checkError.ErrCheck(err)
	println("FRIENDS:")
	for users.Next() {
		var friend Friend
		err = users.Scan(&friend.FriendOneUid, &friend.FriendTwoUid)
		checkError.ErrCheck(err)

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
			checkError.ErrCheck(err)

		}
	}
}

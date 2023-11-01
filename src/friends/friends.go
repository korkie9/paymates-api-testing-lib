package friends

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

func GetUserFriends(db *sql.DB) {
	requestBody := map[string]interface{}{}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("GET", requestBody, "Friends/get-friends", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Friend Response: ", resStr)

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

func AddFriend(db *sql.DB) {
	fmt.Println("Enter the uid of the friend you'd like to add...")
	friendUid, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"friendUid": friendUid,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Friends/add-friend", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Friend Response: ", resStr)
}

func AddMultipleFriends() {
	// fmt.Println("How many mock users would you like to add")
	// //Read input
	// var amount int
	// for {
	// 	input, err := util.GetUserInput()
	// 	check_error.ErrCheck(err)
	// 	//check if input matches an int value
	// 	numberreg := regexp.MustCompile(`\d`)
	// 	if !numberreg.MatchString(input) {
	// 		fmt.Println("Please add a valid number")
	// 	} else {
	// 		castedAmount, err := strconv.Atoi(input)
	// 		check_error.ErrCheck(err)
	// 		amount = castedAmount
	// 		break
	// 	}

	// }
	// //register mock user
	// for i := 0; i < amount; i++ {
	// 	RegisterMockUser(db)
	// }
}

func DeleteFriend(db *sql.DB) {
	fmt.Println("Enter the uid of the friend you'd like to Delete...")
	friendUid, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"FriendUid": friendUid,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("DELETE", requestBody, "Friends/remove-friend", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Delete Friend Response: ", resStr)
}

func TruncateFriends(db *sql.DB) {
	_, err := truncate.Truncate(db, "Friends")
	check_error.ErrCheck(err)
}

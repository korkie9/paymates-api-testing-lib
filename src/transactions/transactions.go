package transactions

import (
	"fmt"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
	util "paymates-mock-db-updater/src/util/get_input"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTransAction() {
	fmt.Println("Enter the uid of the debtor Id...")
	debtorId, err := util.GetUserInput()
	check_error.ErrCheck(err)
	fmt.Println("Enter the uid of the debtor Id...")
	creditorId, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"icon":        "🐐",
		"title":       "Hotdog",
		"amount":      50.50,
		"debtorUid":   debtorId,
		"creditorUid": creditorId,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Transactions/create-transaction", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Create Transaction Response: ", resStr)
}

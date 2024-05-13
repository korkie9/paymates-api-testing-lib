package transactions

import (
	"fmt"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
	util "paymates-mock-db-updater/src/util/get_input"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteAccount() {
	fmt.Println("Enter the uid of the account you'd like to delete...")
	accountId, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"bankAccountUid": accountId,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("DELETE", requestBody, "BankAccounts/remove-bank-account", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Create Transaction Response: ", resStr)
}

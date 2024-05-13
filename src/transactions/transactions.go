package transactions

import (
	"database/sql"
	"fmt"
	"paymates-mock-db-updater/src/auth"
	"paymates-mock-db-updater/src/check_error"
	reqres "paymates-mock-db-updater/src/httpRequest"
	util "paymates-mock-db-updater/src/util/get_input"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Transaction struct {
	Uid               string   `json:"uid"`
	Icon              string   `json:"icon"`
	Title             string   `json:"title"`
	Amount            string   `json:"amount"`
	DebtorUsernames   []string `json:"debtorUsernames"`
	CreditorUsernames []string `json:"creditorUsernames"`
	CreatedAt         string   `json:"createdAt"`
	FriendId          string   `json:"friendId"`
}

func CreateTransAction() {
	fmt.Println("Enter the username of the debtor...")
	debtorInput, err := util.GetUserInput()
	check_error.ErrCheck(err)
	debtorUsernames := strings.Split(debtorInput, " ")
	fmt.Println("Enter the username of the creditor...")
	creditorInput, err := util.GetUserInput()
	check_error.ErrCheck(err)
	creditorUsernames := strings.Split(creditorInput, " ")
	fmt.Println("Enter the amount owed...")
	amount, err := util.GetUserInput()
	check_error.ErrCheck(err)
	fmt.Println("Enter title of transaction...")
	title, err := util.GetUserInput()
	check_error.ErrCheck(err)
	requestBody := map[string]interface{}{
		"title":             title,
		"amount":            amount,
		"debtorUsernames":   debtorUsernames,
		"creditorUsernames": creditorUsernames,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("POST", requestBody, "Transactions/create-transaction", token)
	check_error.ErrCheck(err)
	resStr := string(res)

	fmt.Println("Debtor array length: ", len(debtorUsernames), " and creditor length: ", len(creditorUsernames))
	fmt.Println("Create Transaction Response: ", resStr)
}

func GetAllTransactions(db *sql.DB) {
	users, err := db.Query("SELECT * FROM Transactions")
	check_error.ErrCheck(err)
	println("TRANSACTIONS:")
	for users.Next() {
		var transaction Transaction
		err = users.Scan(&transaction.Uid, &transaction.Icon, &transaction.Title, &transaction.Amount, &transaction.CreatedAt, &transaction.CreditorUsernames, &transaction.DebtorUsernames, &transaction.FriendId)
		check_error.ErrCheck(err)

		fmt.Println(transaction.Amount, ' ', transaction.CreatedAt, " ", transaction.CreditorUsernames, " ", transaction.DebtorUsernames, transaction.FriendId, " ", transaction.Icon, " ", transaction.Title, " ", transaction.Uid)
	}
}

func PrintArray() {
	fmt.Println("Enter the username of the debtor...")
	debtorInput, err := util.GetUserInput()
	check_error.ErrCheck(err)
	debtorUsernames := strings.Split(debtorInput, " ")
	fmt.Println("Slice: ", debtorUsernames)
}

func DeleteTransAction() {
	fmt.Println("Enter the uid of the transaction You'd like to delete...")
	transactionUid, err := util.GetUserInput()
	check_error.ErrCheck(err)

	requestBody := map[string]interface{}{
		"transactionUid": transactionUid,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("DELETE", requestBody, "Transactions/delete-transaction", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Delete Transaction Response: ", resStr)
}

func GetTransaction() {
	fmt.Println("Enter the uid of the transaction You'd like to delete...")
	transactionUid, err := util.GetUserInput()
	check_error.ErrCheck(err)

	requestBody := map[string]interface{}{
		"transactionUid": transactionUid,
	}
	token := auth.GetAccessToken()
	res, err, _ := reqres.HttpRequest("GET", requestBody, "Transactions/delete-transaction", token)
	check_error.ErrCheck(err)
	resStr := string(res)
	fmt.Println("Delete Transaction Response: ", resStr)
}

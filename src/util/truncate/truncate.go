package util

import (
	"database/sql"
	"fmt"
	util "paymates-mock-db-updater/src/util/get_input"
)

func Truncate(db *sql.DB, tableName string) (isDeleted bool, err error) {
	fmt.Println("Do you want to delete all ", tableName, " from the database? Type (yes) if so. Enter any other value if not.")
	input, err := util.GetUserInput()
	if err != nil {
		return false, err
	}
	if input == "yes" {
		_, err = db.Exec("Truncate table " + tableName)
		if err != nil {
			return false, err
		}
		fmt.Println("All ", tableName, " were deleted. Sure hope you know what you're doing.")
		return true, nil
	}
	fmt.Println("Deletion canceled")
	return false, nil
}

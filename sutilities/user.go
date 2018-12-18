package sutilities

import (
	"github.com/Gigamons/Shared/shelpers"
)

func GetUserId(userName string) (error, int32) {
	var UserID int32

	rows, err := shelpers.DB.Query("SELECT id FROM users WHERE UserName=?", userName)
	if err != nil {
		return err, -1
	}
	for rows.Next() {
		err := rows.Scan(&UserID)
		if err != nil {
			return err, -1
		}
	}

	return rows.Close(), UserID
}

func NewUser(userName string, password string) error {
	_, err := shelpers.DB.Exec("INSERT INTO users (UserName, Password) VALUES (?, ?)", userName, shelpers.Generate_Hash(password))
	return err
}

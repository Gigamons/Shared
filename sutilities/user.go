package sutilities

import (
	"github.com/cyanidee/bancho-go/helpers"
	"strings"
)

func GetUserId(userName string) (error, int32) {
	var UserID int32
	lower := strings.ToLower(strings.Replace(userName, " ", "_", -1))

	rows, err := helpers.DBConn.Query("SELECT id FROM users WHERE username_lower=?", lower)
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

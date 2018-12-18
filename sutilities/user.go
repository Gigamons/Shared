package sutilities

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Gigamons/Shared/shelpers"
)

type User struct {
	Id                    int
	UserName              string
	EMail                 string
	Password              string
	Privileges            int
	Achievements          int
	AchievementsDisplayed int
}

func GetUserId(userName string) (int32, error) {
	var UserID int32

	rows, err := shelpers.DB.Query("SELECT id FROM users WHERE UserName=?", userName)
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		err := rows.Scan(&UserID)
		if err != nil {
			return -1, err
		}
	}

	return UserID, rows.Close()
}

func NewUser(userName string, password string) error {
	m := md5.New()
	m.Write([]byte(password))
	_, err := shelpers.DB.Exec("INSERT INTO users (UserName, Password) VALUES (?, ?)", userName, shelpers.Generate_Hash(hex.EncodeToString(m.Sum(nil))))
	return err
}

func GetUser(userId int32) (*User, error) {
	u := new(User)

	rows, err := shelpers.DB.Query("SELECT id, UserName, EMail, Password, Privileges, Achievements, AchievementsDisplayed FROM users WHERE id=?", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.UserName, &u.EMail, &u.Password, &u.Privileges, &u.Achievements, &u.AchievementsDisplayed)
		if err != nil {
			return nil, err
		}
	}

	return u, rows.Close()
}

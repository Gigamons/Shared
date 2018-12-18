package shelpers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func AntiTimeout(db *sql.DB) {
	for {
		if err := db.Ping(); err != nil {
			log.Fatalln(err)
			break
		}

		time.Sleep(time.Second * 10)
	}
}

func ConnectMySQL(hostname string, port uint16,
	         username string, password string,
	         database string) (db *sql.DB, err error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)
	db, err = sql.Open("mysql", conStr)

	DB = db

	go AntiTimeout(DB)
	return
}
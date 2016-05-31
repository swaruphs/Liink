package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Print("db error %v", err.Error())
	}

	if err = DB.Ping(); err != nil {
		//log.Panic(err)
		fmt.Print(err)
	}

	createTables(DB)
}

func createTables(db *sql.DB) (bool, error) {

	_, err := db.Exec("create table if not exists link (id serial primary key, weburl varchar(200) not null, appstore_url varchar(200) not null, playstore_url varchar(200) not null)")
	if err != nil {
		return false, err
	}

	return true, nil
}

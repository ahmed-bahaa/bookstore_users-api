package useres_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_userdb_username = "mysql_userdb_username"
	mysql_userdb_password = "mysql_userdb_password"
	mysql_userdb_host     = "mysql_userdb_host"
	mysql_userdb_schema   = "mysql_userdb_schema"
)

var (
	Client   *sql.DB
	err      error
	username = os.Getenv(mysql_userdb_username)
	password = os.Getenv(mysql_userdb_password)
	host     = os.Getenv(mysql_userdb_host)
	schema   = os.Getenv(mysql_userdb_schema)
)

func init() {
	//?charset=utf8
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = Client.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("database successfully configured!")
}

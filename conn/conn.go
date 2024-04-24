package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(host string, port string, user string, password string, dbname string) {
	// ...
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	// 尝试ping数据库以验证连接是否成功
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Success")
	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
}

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(host string, port string, username string, password string, dbname string) bool {
	// ...
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	// 尝试ping数据库以验证连接是否成功
	err = conn.Ping()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	fmt.Println("Connection Success")
	return true
}

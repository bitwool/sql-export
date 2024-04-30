package db

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var TableName string

func Init(host string, port string, user string, password string, dbname string, query string) {
	connect(host, port, user, password, dbname)
	parseTableName(query)
}

func connect(host string, port string, user string, password string, dbname string) {
	// ...
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	// 尝试ping数据库以验证连接是否成功
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	DB = conn
	fmt.Println("Connection Success")
}

func Close() {
	DB.Close()
}

func GetAipk(query string) string {

	// 查询所有字段信息
	queryColumns := `
		SELECT column_name, extra
		FROM information_schema.columns
		WHERE table_name = ? AND table_schema = database()
	`
	columns := make(map[string]string)
	rows, err := DB.Query(queryColumns, TableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var columnName, extra string
		err := rows.Scan(&columnName, &extra)
		if err != nil {
			log.Fatal(err)
		}
		columns[columnName] = extra
	}

	// 查询主键信息
	queryPrimaryKey := `
		SELECT column_name
		FROM information_schema.table_constraints AS tc
		INNER JOIN information_schema.key_column_usage AS ku
		ON tc.constraint_name = ku.constraint_name
		WHERE tc.constraint_type = 'PRIMARY KEY' AND ku.table_name = ? AND ku.table_schema = database()
	`
	primaryKey := ""
	rows, err = DB.Query(queryPrimaryKey, TableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&primaryKey)
		if err != nil {
			log.Fatal(err)
		}
		break // 假设每个表只有一个主键
	}

	// 检查自增且为主键的列
	if extra, ok := columns[primaryKey]; ok && extra == "auto_increment" {
		fmt.Printf("Column '%s' is an auto-incrementing primary key.\n", primaryKey)
		return primaryKey
	} else {
		return ""
	}
}

func parseTableName(query string) {

	// 使用正则表达式匹配表名
	re := regexp.MustCompile(`FROM\s+([a-zA-Z0-9_]+)`)
	matches := re.FindStringSubmatch(query)

	// 提取表名
	TableName = matches[1]
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

type Item struct {
	DbType string `json:"dbType"`
	Field  string `json:"filed"`
	Value  string `json:"value"`
}

type Items struct {
	Items []Item `json:"items"`
}

func (a *App) Query(host string, port string, username string, password string, dbname string, query string) []Items {

	result := []Items{}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error:", err)
		return result
	}
	// 尝试ping数据库以验证连接是否成功
	err = conn.Ping()
	if err != nil {
		fmt.Println("Error:", err)
		return result
	}

	fmt.Println("Connection Success")
	fmt.Println(query)

	rows, err := conn.Query(query)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return result
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		newItems := Items{
			Items: make([]Item, len(columns)),
		}
		for i := range columns {
			newItems.Items[i] = Item{
				DbType: "string",
				Field:  columns[i],
				Value:  string(values[i]),
			}
		}
		result = append(result, newItems)
	}

	fmt.Println(result)

	return result
}

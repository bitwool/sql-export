package export

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/bitwool/sql-export/db"

	_ "github.com/go-sql-driver/mysql"
)

func Query(query string, ignore_ai_pk bool) {

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// columnTypes, err := rows.ColumnTypes()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if ignore_ai_pk {
		aipk := db.GetAipk(query)
		fmt.Print("aipk:" + aipk)
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	output := "INSERT INTO " + db.TableName + " (" + strings.Join(columns, ",") + ") VALUES "

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		strData := make([]string, len(values))
		for i, rawBytes := range values {
			strData[i] = "'" + string(rawBytes) + "'" // 转换为字符串
		}

		output += "(" + strings.Join(strData, ",") + "),"
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output[0:len(output)-1] + ";")
}

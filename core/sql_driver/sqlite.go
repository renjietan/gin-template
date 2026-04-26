package sql_driver

import (
	"fmt"

	"example.com/t/types"
	"example.com/t/utility"
)

/**
 * @FILE   sqlite
 * @AUTHOR TAN
 * @DESCRIPTION
 * @DATE 14:02:43 CST 2026-04-24
 * @PARAM
 * @RETURN
 **/
func NewSqliteDriver(appConfig *types.AppConfig) (string, error) {
	dbFile := utility.Tern(appConfig.Debug == true, "sqlite3_dev.db", "sqlite3_prod.db")
	fmt.Println("dbFile", dbFile)
	return "dbFile", nil
}

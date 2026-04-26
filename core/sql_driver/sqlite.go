package sql_driver

import (
	"fmt"
	"path/filepath"

	"example.com/t/types"
	"example.com/t/utility"
	sqlcipher "github.com/gdanko/gorm-sqlcipher"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
 * @FILE   sqlite
* @AUTHOR TAN
 * @DESCRIPTION
 * @DATE 14:02:43 CST 2026-04-24
 * @PARAM
 * @RETURN
 **/
func NewSqliteDriver(config *gorm.Config, appConfig *types.AppConfig) (*gorm.DB, error) {
	//_pragma_cipher_page_size=4096
	params := fmt.Sprintf("?_pragma_key=x'%s'&_pragma_cipher_compatibility=3", appConfig.Sqlite.SqliteEncryptionKey)
	dbFileName := utility.Tern(appConfig.Debug == false, "dev.db", "prod.db"+params)
	dbFullPath := filepath.Join(appConfig.Sqlite.BasePath, appConfig.AppName+"-"+dbFileName.(string))
	//db, err := gorm.Open(sqlite.Open(dns), config)
	//db, err := gorm.Open(sqlite.Open(dbFullPath), config)
	db, err := gorm.Open(sqlcipher.Open(dbFullPath), config)
	if err != nil {
		return nil, err
	}
	return db, nil
}

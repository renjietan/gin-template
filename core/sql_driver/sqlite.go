package sql_driver

import (
	"fmt"
	"path/filepath"

	"example.com/t/types"
	"example.com/t/utility"
	//sqlcipher "github.com/gdanko/gorm-sqlcipher"
	sqliteEncrypt "github.com/ShaoQ1ang/gorm-sqlite-cipher"
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
type User struct {
	ID   uint
	Name string
}

func NewSqliteDriver(config *gorm.Config, appConfig *types.AppConfig) (*gorm.DB, error) {
	// _pragma_cipher_page_size=4096
	// _pragma_cipher_compatibility=3
	params := fmt.Sprintf("?_pragma_key=%s&_pragma_cipher_page_size=4096", appConfig.Sqlite.SqliteEncryptionKey)
	dbFileName := utility.Tern(appConfig.Debug, "dev.db", "prod.db"+params)
	dbFullPath := filepath.Join(appConfig.Sqlite.BasePath, appConfig.AppName+"-"+dbFileName.(string))
	//db, err := gorm.Open(sqlite.Open(dbFullPath), config)
	db, err := gorm.Open(sqliteEncrypt.Open(dbFullPath), config)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Alice"})
	var user User
	db.First(&user)
	println("查询结果:", user.Name)
	return db, nil
}

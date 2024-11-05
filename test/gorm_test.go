package test

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestCumulus(t *testing.T) {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/demo")
	if err != nil {
		return
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	types, err := gormDB.Migrator().ColumnTypes("users")
	if err != nil {
		return
	}

	for _, columnType := range types {
		fmt.Println(columnType.Name())
		fmt.Println(columnType.DefaultValue())
	}
}

func TestStrings(t *testing.T) {

	fmt.Println(strings.Repeat("aaa", 2))

}

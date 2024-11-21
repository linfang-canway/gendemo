package dal

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
		// DB = DB.Debug()
	})
}

func ConnectDB() (conn *gorm.DB) {
	var err error

	conn, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}

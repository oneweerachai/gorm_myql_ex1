package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/users_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	var now = time.Now()
	var age = 17
	user := User{Name: "Jinzhu", Age: age, Birthday: &now}

	result := db.Create(&user) // pass pointer of data to Create

	//user.ID // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	fmt.Printf("user.ID = %v  rows affected: %v", user.ID, result.RowsAffected)

	// db.Model(&User{}).Create(map[string]interface{}{
	// 	"Name": "Weera", "Age": 30,
	// })

}

type User struct {
	ID           int
	Name         string
	Email        *string
	Age          int
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

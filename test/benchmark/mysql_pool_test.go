package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	ID   uint
	Name string
}
func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Tucker "}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}
func BenchmarkMaxOpenConns(b *testing.B) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_backend?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false,
		// Logger: false,
	})
	if err != nil {
		log.Fatalln("fail to connect database: ", err)
	}
	// check if table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exits
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// handle error if you want
			fmt.Println("error dropping tabler: ", err)
		}
	}
	// create table if table does not exists
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("failed to get sql.DB from gorm.DB: ", err)
	}
	// set connection arguments
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
// Benchmark in Go is the process of measuring code performance to determine execution speed, resource usage, or optimize code, using the testing package with functions prefixed by Benchmark.
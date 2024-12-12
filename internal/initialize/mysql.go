package initialize

import (
	"fmt"
	"time"

	"github.com/DoCongThanhPhuong/go-backend/global"
	"github.com/DoCongThanhPhuong/go-backend/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func logs the errors
func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
// func initializes the MySQL connection
func InitMysql() {
	m := global.Config.MySQL
	// build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.DBName)
	// open the connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to initialize MySQL")
	global.Logger.Info("MySQL Initialized successfully")
	global.Mdb = db
	// set connection pool settings
	// a pool is a set of pre-maintained connections that improve performance.
	setPool()
	// run migrations
	migrateTables()
}
// func set the MySQL connection pool settings
func setPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	checkErrorPanic(err, "Failed to get Database from GORM")
	// set connection pool configurations
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

// func runs database migrations
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("Error during table migration", zap.Error(err))
	}
}
package shared_infra_db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *sql.DB {
	// 	config, err := shared_configs.GetConfig()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	dsn := fmt.Sprintf("%s:%s@/%s",
	// 		config.Database.DB.Username,
	// 		config.Database.DB.Password,
	// 		config.Database.DB.DBName)
	// 	db, err := sql.Open(config.Database.Driver, dsn)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	db.SetConnMaxLifetime(time.Second * 3)
	// 	db.SetMaxOpenConns(10)
	// 	db.SetMaxIdleConns(10)
	// 	return db
	return nil
}

func NewMysqlDB_GORM() *gorm.DB {
	// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 		config.Database.DB.Username,
	// 		config.Database.DB.Password,
	// 		config.Database.DB.Host,
	// 		config.Database.DB.Port,
	// 		config.Database.DB.DBName)

	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	return db
	return nil
}

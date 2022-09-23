// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 15:49$

package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"upserver/internal/pkg/utils"
)

var (
	db *gorm.DB
)

const (
	DriverPostgresql = "postgresql"
	DriverMysql      = "mysql"
	DriverMongo      = "mongodb"
)

// InitDB
// @description: 初始化数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:37
// @success:
func InitDB() {
	var err error
	var dsn = MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", utils.Config.Mysql.DBName)
	// 创建数据库
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库创建失败")
		return
	}

	//数据库驱动
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Config.Mysql.UserName,
		utils.Config.Mysql.Password,
		utils.Config.Mysql.Host,
		utils.Config.Mysql.Port,
		utils.Config.Mysql.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
	})
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(utils.Config.Mysql.MinConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(utils.Config.Mysql.MaxConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Info("init db success")

}

// GetDB
// @description: 获取数据库连接
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:38
// @success:
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	//fmt.Println(fmt.Sprintf("%#v", db))
	return db
}

// createDatabase 创建数据库（ EnsureDB() 中调用 ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func MysqlEmptyDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", utils.Config.Mysql.UserName,
		utils.Config.Mysql.Password,
		utils.Config.Mysql.Host,
		utils.Config.Mysql.Port)
}

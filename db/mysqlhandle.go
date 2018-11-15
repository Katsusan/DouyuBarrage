package db

import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	""
)

var MainDB *xorm.Engine

var (
	ConnectErr  = errors.New("Can't connect to mysql")
	CreateDBErr = errors.New("Can't create specified database")
)

func initEngine() error {
	var err error

	username := "root"
	password := "katsu2017"
	dbname := "dybarrage"
	charset := "utf8"
	conninfo := username + ":" + password + "@/" + dbname + "?charset=" + charset
	maxIdle := 2
	maxConn := 10

	MainDB, err := xorm.NewEngine("mysql", conninfo)
	if err != nil {
		return err
	}

	MainDB.SetMaxIdleConns(maxIdle)
	MainDB.SetMaxOpenConns(maxConn)

	//检测DB连接是否正常
	if err = MainDB.Ping(); err != nil {
		return ConnectErr
	}

	//指定使用的Database
	_, err = MainDB.Exec("use " + dbname)
	if err != nil {
		log.Printf("use DB(%s) error.\n", dbname)
		_, err = MainDB.Exec("CREATE DATABASE " + dbname + " DEFAULT CHARACTER SET " + charset)
		if err != nil {
			log.Printf("create DB(%s) error.\n", dbname)
			return CreateDBErr
		}
		log.Printf("succeed to create database(%s).\n", dbname)
	}

	//检测表是否存在
	tbexist, err := MainDB.IsTableExist(new(data.Barrage))

	return nil

}

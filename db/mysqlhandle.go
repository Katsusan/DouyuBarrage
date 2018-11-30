package db

import (
	"errors"
	"log"

	"../data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var MainDB = new(xorm.Engine)

var (
	ErrConnect  = errors.New("Can't connect to mysql")
	ErrCreateDB = errors.New("Can't create specified database")
	ErrTableChk = errors.New("Can't check the existence of table")
	ErrTableCrt = errors.New("Can't create necessary table")
	ERRTableSyn = errors.New("Can't synchronize structs to table")
)

func init() {
	log.Printf("db.init exec.\n")
}

func InitEngine(dbcfg *data.MysqlConfig) error {
	log.Printf("db.InitEngine exec.\n")
	var err error

	MainDB, err = xorm.NewEngine("mysql", dbcfg.ConnectInfo())
	//log.Printf("coninfo:%s\n", dbcfg.ConnectInfo())
	if err != nil {
		return err
	}

	MainDB.SetMaxIdleConns(2)
	MainDB.SetMaxOpenConns(10)

	//检测DB连接是否正常
	if err = MainDB.Ping(); err != nil {
		return ErrConnect
	}

	//指定使用的Database
	_, err = MainDB.Exec("use " + dbcfg.DbName)
	if err != nil {
		log.Printf("use DB(%s) error.\n", dbcfg.DbName)
		_, err = MainDB.Exec("CREATE DATABASE " + dbcfg.DbName + " DEFAULT CHARACTER SET " + dbcfg.CharSet)
		if err != nil {
			log.Printf("create DB(%s) error.\n", dbcfg.DbName)
			return ErrCreateDB
		}
		log.Printf("succeed to create database(%s).\n", dbcfg.DbName)
	}

	//检测表(barrage_<roomid>)是否存在
	tmpdt := new(data.Barrage)
	tmpdt.Roomid = dbcfg.TableID
	tbexist, err := MainDB.IsTableExist(tmpdt)
	if err != nil {
		log.Println("check table existence failed.")
		return ErrTableChk
	}
	if !tbexist {
		log.Println("table not exist, need to create a new table.")
		err = MainDB.CreateTables(tmpdt)
		if err != nil {
			log.Println("create table failed.")
			return ErrTableCrt
		}
		return nil
	}

	//已存在表的话同步表数据结构
	err = MainDB.Sync2(tmpdt)
	if err != nil {
		log.Println("table exists, but sync failed.")
		return ERRTableSyn
	}

	return nil

}

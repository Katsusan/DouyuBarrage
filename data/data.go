package data

import (
	"fmt"
	"time"
)

const (
	TABLE_PREFIX = "barrage_"
)

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	CharSet  string
	TableID  string
}

func (cfg *MysqlConfig) ConnectInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.CharSet)
}

type Barrage struct {
	Id        int64
	Nickname  string    `xorm:"varchar(255) notnull 'nickname'"`
	Chatmsg   string    `xorm:"varchar(512) notnull 'chatmsg'"`
	Roomid    string    `xorm:"int notnull 'roomid'"`
	CreatedAt time.Time `xorm:"created"`
}

func (br *Barrage) TableName() string {
	return TABLE_PREFIX + br.Roomid
}

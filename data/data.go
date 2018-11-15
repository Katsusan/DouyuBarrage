package data

import "time"

type Barrage struct {
	Id        int64
	Nickname  string    `xorm:"varchar(255) notnull 'nickname'"`
	Chatmsg   string    `xorm:"varchar(512) notnull 'chatmsg'"`
	CreatedAt time.Time `xorm:"created"`
}

func (br *Barrage) TableName() string {
	return "barrage"
}

package shortlink_repo

import "time"

// Short4LongLinkModel 短链数据库model
type Short4LongLinkModel struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	ShortLink string    `xorm:"index" json:"shortLink"`
	LongLink  string    `json:"longLink"`
	Ctime     time.Time `xorm:"created" json:"ctime"`
	Utime     time.Time `xorm:"updated" json:"utime"`
	IsDelete  byte      `json:"isDelete"`
}

func (Short4LongLinkModel) TableName() string {
	return "t_short_long_link"
}

package ws

import "time"

type Trainer struct {
	Content   string    `db:"content"` //内容
	StartTime time.Time `db:"time"`    //创建时间
	Read      uint      `db:"read"`    //已读
}

type Result struct {
	StartTime time.Time
	Msg       string
	Content   interface{}
	From      string
}

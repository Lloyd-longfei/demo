package model

type Test struct {
	Id        int    `json:"id" xorm:"pk autoincr"`
	Name      string `json:"name" xorm:"varchar(20) notnull"`
	OrderId   string `json:"order_id" xorm:"varchar(30) not null"`
	CreatedAt int64  `json:"created_at" xorm:"created"`
	UpdatedAt int64  `json:"updated_at" xorm:"updated"`
}

package model

type User struct {
	Id          int    `gorm:"id" json:"id"`
	Uid         int    `gorm:"uid" json:"uid"`
	Name        string `gorm:"name" json:"name"`
	Phone       string `gorm:"phone" json:"phone"`
	Avatar      string `gorm:"avatar" json:"avatar"`
	Status      int    `gorm:"status" json:"status"`
	CreatedTime string `gorm:"created_time" json:"created_time"`
	UptadeTime  string `gorm:"uptade_time" json:"uptade_time"`
}

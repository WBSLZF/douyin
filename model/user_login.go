package model

type UserLogin struct {
	Id         int64
	UserCount  string `gorm:"notnull"`
	PassWord   string `gorm:"notnull"`
	UserInfoId int64  //账户密码的一对一关系
}

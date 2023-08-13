package model

type UserLogin struct {
	Id         int
	UserCount  string `gorm:"notnull"`
	PassWord   string `gorm:"notnull"`
	UserInfoId int    //账户密码的一对一关系
}

package model

type User struct {
	Id       int64  `description:"id" json:"id" xorm:"pk autoincr"`
	NickName string `description:"昵称 最长50字符" json:"nickName" valid:"MaxSize(50)"`
	UserName string `description:"登录名" json:"userName"`
	Password string `description:"密码,加密存储"`
	Created  int64  `description:"注册时间" xorm:"created"`
	Updated  int64  `description:"修改时间" xorm:"updated"`
	Deleted  int64  `description:"删除时间" json:"-" xorm:"deleted"`
	UserType int 	`description:"用户类型,1为超级管理员，0为普通用户，-1为停用用户" json:"userType"`

}

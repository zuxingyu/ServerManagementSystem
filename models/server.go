package models

type Server struct {
	Id          int64  `description:"id" json:"id" xorm:"pk autoincr"`
	AliasName 	string `description:"别名" json:"aliasName"`
	IpAddr 	string `description:"IP地址" json:"ipAddr"`
	IpPort	string `description:"IP端口" json:"ipPort"`
	Account string `description:"用户名" json:"account"`
	Password string `description:"密码" json:"password"`
	UserId int64 `description:"用户编号" json:"userId"`
	ServerOs string `description:"操作系统" json:"serverOS"`
	CpuType string `description:"CPU型号" json:"CPUType"`
	CpuNum int `description:"CPU数量" json:"CPUNum"`
	MemoryTotal int64 `description:"内存空间" json:"memoryTotal"`
	HardDisk int64 `description:"硬盘空间" json:"hardDisk"`
}
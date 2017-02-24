package models

type Server struct {
	Id          int64  `description:"id" json:"id" xorm:"pk autoincr"`
	AliasName   string `description:"别名" json:"aliasName"`
	IpAddr      string `description:"IP地址" json:"ipAddr"`
	IpPort      string `description:"IP端口" json:"ipPort"`
	Account     string `description:"用户名" json:"account"`
	Password    string `description:"密码" json:"password"`
	UserId      int64  `description:"用户编号" json:"userId"`
	ServerOS    string `description:"操作系统" json:"serverOS"`
	CPUType     string `description:"CPU型号" json:"CPUType"`
	CPUNum      int    `description:"CPU数量" json:"CPUNum"`
	MemoryTotal int64  `description:"内存空间" json:"memoryTotal"`
	HardDisk    int64  `description:"硬盘空间" json:"hardDisk"`
	ServerState int64  `descriptipn:"服务器状态,0为正常，1为停用" json:"serverState"`
	Created     int64  `description:"注册时间" xorm:"created"`
	Updated     int64  `description:"修改时间" xorm:"updated"`
}

type ServerDTO struct {
	Id          int64   `description:"id" json:"id" xorm:"pk autoincr"`
	AliasName   string  `description:"别名" json:"aliasName"`
	ServerOS    string  `description:"操作系统" json:"serverOS"`
	ServerState int64   `descriptipn:"服务器状态" json:"serverState"`
	CPURate     float64 `descriptipn:"CPU负载状态" json:"cpuRate"`
	MemoryRate  float64 `descriptipn:"内存负载状态" json:"memoryRate"`
	Created     int64   `description:"注册时间" xorm:"created"`
}


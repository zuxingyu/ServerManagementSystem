package models

type ServerMemoryLog struct {
	Id      	int64   `description:"id" json:"id" xorm:"pk autoincr"`
	Total   	int64   `description:"物理内存总量" json:"total"`
	Used    	int64   `description:"使用的物理内存总量" json:"used"`
	Free    	int64   `description:"空闲内存总量" json:"free"`
	Buffers 	int64   `description:"缓存内存总量" json:"buffers"`
	MemoryRate 	float64 `description:"内存使用比" json:"memoryRate"`
	Cached		int64   `description:"缓冲的交换区总量" json:"cached"`
	ServerId 	int64   `description:"服务器编号" json:"serverId"`
	Created     	int64   `description:"创建时间" xorm:"created"`
}
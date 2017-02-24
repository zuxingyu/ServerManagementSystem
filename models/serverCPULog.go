package models

type ServerCPULog struct {
	Id       int64      `description:"id" json:"id" xorm:"pk autoincr"`
	CPURate  float64    `description:"用户空间占用CPU百分比" json:"CPUrate"`
	SyRate   float64    `description:"内核空间占用CPU百分比" json:"syRate"`
	NiRate   float64    `description:"用户进程空间占用CPU百分比" json:"niRate"`
	IdRate   float64    `description:"空闲CPU百分比" json:"idRate"`
	WaRate   float64    `description:"等待输出CPU百分比" json:"waRate"`
	HiRate   float64    `description:"硬中断CPU百分比" json:"hiRate"`
	SiRate   float64    `description:"软中断CPU百分比" json:"siRate"`
	ServerId int64      `description:"服务器编号" json:"serverId"`
	Created  int64      `description:"创建时间" xorm:"created"`
}

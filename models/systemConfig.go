package models

type SystemConfig struct {
	Id          int64  `description:"id" json:"id" xorm:"pk autoincr"`
	Param       string `description:"参数" json:"param"`
	Value       string `description:"值" json:"value"`
	Instruction string `description:"参数说明" json:"instruction"`
	Timestamp   int64  `description:"时间戳" json:"timestamp"`
}

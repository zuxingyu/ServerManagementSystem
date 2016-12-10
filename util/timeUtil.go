package util

import "time"


//获取当前时间
func GetCurrentDate_YYYYMMDD()  {

}

func BeijingTime() time.Time {
	localTime := time.Now()
	location, _ := time.LoadLocation("Asia/Shanghai")
	beijingTime := localTime.In(location)
	return beijingTime
}

func UnixOfBeijingTime() int64 {
	localTime := time.Now()
	location, _ := time.LoadLocation("Asia/Shanghai")
	beijingTime := localTime.In(location)
	return beijingTime.Unix()
}

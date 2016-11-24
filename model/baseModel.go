package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/go-xorm/xorm"
	"fmt"
	"os"
	"time"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"ServerManagementSystem/logs"
)

const (
	RUN_MODE_DEV = "dev"
	RUN_MODE_TEST = "test"
	RUN_MODE_PROD = "prod"
)


type DBConfig struct {
	DbType   string
	User     string
	Password string
	Database string
	Charset  string
	Host     string
	Port     string
}

var redisConn string
var Orm *xorm.Engine
var RedisCache cache.Cache

func init(){
	var dbConfig = DBConfig{}
	if beego.BConfig.RunMode == RUN_MODE_DEV {
		dbConfig.DbType = "mysql"
		dbConfig.User = "root"
		dbConfig.Password = "root"
		dbConfig.Database = "SMS"
		dbConfig.Charset = "utf8mb4"
		dbConfig.Host = "localhost"
		dbConfig.Port = ":3306"
		redisConn = "localhost:6379"
	}
	InitRedis(`{"key":"smsRedies","conn":"` + redisConn + `","dbNum":"1"}`)
	InitOrm(dbConfig)
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	err := Orm.Sync2(new(User))

	if err != nil {
		logs.Logger.Error("同步数据结构失败！\n" + err.Error())
		panic("同步数据结构失败!\n" + err.Error())
	}


}

func InitRedis(redisConfig string) {
	var err error
	RedisCache, err = cache.NewCache("redis", redisConfig)
	if err != nil {
		panic(err)
	}
}

func InitOrm(dbConfig DBConfig){
	var err error
	Orm, err = xorm.NewEngine(dbConfig.DbType, fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset))
	if err != nil {
		panic("创建数据库连接Engine失败!")
	}
	Orm.ShowSQL(true)
	Orm.SetMaxIdleConns(50)
	Orm.SetMaxOpenConns(500)
	f, err := os.Create("data/log/sql.log")
	if err != nil {
		logs.Logger.Error("创建sql.log文件失败！")
	}
	// defer f.Close()
	Orm.SetLogger(xorm.NewSimpleLogger(f))
}
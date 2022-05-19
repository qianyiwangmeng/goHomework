/**
  @author: qianyi  2022/5/17 20:35:00
  @note:
*/
package main

import (
	"goHomework/config"
	"goHomework/db"
	"goHomework/router"
)

func main() {
	err := config.InitJson("./config.json")
	if err != nil {
		panic(err)
	}
	err = db.InitMysql(config.Conf.MysqlConfig)
	if err != nil {
		panic(err)
	}
	err = db.InitRedis(config.Conf.RedisConfig)
	if err != nil {
		panic(err)
	}
	engine := router.InitRouter()
	engine.Run()
}

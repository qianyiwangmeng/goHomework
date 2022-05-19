/**
  @author: qianyi  2022/5/14 17:09:00
  @note:
*/
package db

import (
	"fmt"
	"goHomework/config"
	"goHomework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(cfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
	err = DB.AutoMigrate(&model.User{}, &model.Article{}, model.Album{})
	if err != nil {
		return err
	}
	return nil
}

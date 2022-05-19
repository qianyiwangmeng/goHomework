/**
  @author: qianyi  2022/5/14 16:18:00
  @note:	用户登录的数据操作
*/

package dao

import (
	"goHomework/db"
	"goHomework/model"
)

// 根据name查询用户
func FindUser(name string) (*model.User, error) {
	var data model.User
	err := db.DB.Model(&model.User{}).Where("name=?", name).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// 创建一个用户
func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

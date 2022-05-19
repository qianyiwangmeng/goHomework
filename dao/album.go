/**
  @author: qianyi  2022/5/17 21:10:00
  @note:
*/
package dao

import (
	"goHomework/db"
	"goHomework/model"
)

// 增加图片
func AddAlbum(album *model.Album) error {
	return db.DB.Create(&album).Error
}

// 展示图片
func ShowAlbum(loginUser string) (*[]model.Album, error) {
	var albums []model.Album
	err := db.DB.Model(&model.Album{}).Where("create_user=?", loginUser).Find(&albums).Error
	if err != nil {
		return nil, err
	}
	return &albums, nil
}

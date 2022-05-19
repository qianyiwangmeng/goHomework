/**
  @author: qianyi  2022/5/17 20:09:00
  @note:
*/
package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"autoIncrement"` // 默认情况下，名为 `ID` 的字段会作为表的主键
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Article struct {
	ID        int       `json:"id" gorm:"autoIncrement"`
	Title     string    `json:"title"`
	Tags      string    `json:"tags"`
	Short     string    `json:"short"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Album struct {
	ID         int       `json:"id" gorm:"autoIncrement"`
	Filepath   string    `json:"filepath"`
	Filename   string    `json:"filename"`
	CreateUser string    `json:"createUser"`
	CreatedAt  time.Time `json:"createAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

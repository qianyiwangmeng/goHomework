/**
  @author: qianyi  2022/5/17 21:10:00
  @note:
*/
package dao

import (
	"goHomework/db"
	"goHomework/model"
	"gorm.io/gorm/clause"
	"log"
)

// 分页获取所有的文章
func LoadArticles(pageIndex, pageSize int) ([]*model.Article, error) {
	var articles []*model.Article
	err := db.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&articles).Error
	if err != nil {
		log.Printf("加载文章失败")
		return nil, err
	}
	return articles, err
}

// 写文章
func AddArticles(article *model.Article) error {
	return db.DB.Create(article).Error
}

// 删除文章
func DeleteArticle(id int) error {
	return db.DB.Delete(&model.Article{}, id).Error
}

// 改文章
func UpdateArticles(article *model.Article) error {
	return db.DB.Where("id=?", article.ID).Updates(article).Error
}

// 根据id查看文章
func FindArticle(id int) (*model.Article, error) {
	var article model.Article
	err := db.DB.Where("id=?", id).Find(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 查询热点文章的信息
func FindTopNArticle(ids []int64) ([]*model.Article, error) {
	var articles []*model.Article

	err := db.DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Find(&articles, ids).Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

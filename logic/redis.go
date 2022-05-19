/**
  @author: qianyi  2022/5/18 20:10:00
  @note:
*/
package logic

import (
	"fmt"
	"goHomework/dao"
	"goHomework/db"
	"goHomework/model"
	"log"
	"strconv"
	"time"
)

const (
	KeyArticleCount = "blog:article:read:count:%s" // 24小时文章阅读数key eq:blog:article:count:20200315
)

// 点击一下加一的操作
func IncArticleReadCount(id string) error {

	todayStr := time.Now().Format("20060102")
	key := fmt.Sprintf(KeyArticleCount, todayStr)

	return db.Client.ZIncrBy(key, 1, id).Err()
}

// 获取热点文章的操作
func GetArticleReadCountTopN(n int64) ([]*model.Article, error) {

	todayStr := time.Now().Format("20060102")
	key := fmt.Sprintf(KeyArticleCount, todayStr)

	idStrs, err := db.Client.ZRevRange(key, 0, n-1).Result()

	log.Printf("Top:%d,redis数据%s", n, idStrs)
	if err != nil {
		return nil, err
	}

	ids := []int64{}

	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("获取热点文章的操作数据转换出错%v", err)
			continue
		}

		ids = append(ids, id)
	}

	articles, err := dao.FindTopNArticle(ids)
	if err != nil {
		log.Printf("获取热点文章失败%v", err)
		return nil, err
	}

	return articles, nil

}

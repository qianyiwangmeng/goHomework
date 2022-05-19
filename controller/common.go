/**
  @author: qianyi  2022/5/17 22:02:00
  @note:
*/
package controller

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ReqArticle struct {
	Title   string `json:"title"`
	Tags    string `json:"tags"`
	Short   string `json:"short"`
	Content string `json:"content"`
}

type ReqUpdateArticle struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Tags    string `json:"tags"`
	Short   string `json:"short"`
	Content string `json:"content"`
}

type ReqPage struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

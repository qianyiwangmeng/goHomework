/**
  @author: qianyi  2022/5/17 22:03:00
  @note:
*/
package utils

import (
	"crypto/md5"
	"fmt"
)

const (
	secret = "轻轻地我将离开你"
)

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum(append([]byte(str), []byte(secret)...)))
	return md5str
}

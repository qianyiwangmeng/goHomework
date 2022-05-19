/**
  @author: qianyi  2022/5/17 20:22:00
  @note:
*/
package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Conf *AppConfig
)

type AppConfig struct {
	*MysqlConfig `json:"mysql"`
	*RedisConfig `json:"redis"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
}

// redis配置
type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
}

func InitJson(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("读取配置文件失败：%v\n", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Conf)

	if err != nil {
		log.Println("Error:", err)
	}
	return nil
}

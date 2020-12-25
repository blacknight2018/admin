package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	DbConnect string `yaml:"db_connect"`
	Port      int    `yaml:"port"`
	SecretKey string `yaml:"secretKey"`
}

var c conf

func GetConf() conf {
	yamlFile, err := ioutil.ReadFile("Config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func GetBindPort() int {
	return c.Port
}

func GetJWTSecret() string {
	return c.SecretKey
}
func GetOneDB() *gorm.DB {
	db, err := gorm.Open("mysql", c.DbConnect)
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxIdleTime(5)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	return db
}

package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"user/model"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// Init 加载配置文件，并返回连接字符串；
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
package main

import (
	"blogBackend/dao/mysql"
	"blogBackend/logger"
	"blogBackend/routers"
	"blogBackend/settings"
	"fmt"
)

// 220301：所谓帖子暂时其实只是文本信息；
// 计划：使用 md 作为帖子本体
// 220302：暂未实现用户系统，博客的前后台是两个VUE项目，后台可用 jwt 验证；
func main() {
	// 初始化配置；
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 初始化日志记录；
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 初始化 mysql 连接；
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 建立路由；
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

package user_test

import (
	"testing"

	"example/internal/global"
	"example/internal/initialize/config"
	"example/internal/initialize/logger"
	"example/internal/initialize/mysql"
	"example/internal/services/user"
)

var userSvc *user.User

func init() {
	// 1、初始化本地配置文件
	config.InitAuto(global.ConfigFile)

	// 2、初始化日志
	logger.Init(false, logger.Hooks())

	// 3、初始化MySQL
	mysql.Init()

	// 4、实例化user服务
	userSvc = user.New()
}

func TestGetCompanyName(t *testing.T) {
	res, err := userSvc.GetCompanyName(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

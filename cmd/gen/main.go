package main

import (
	"flag"
	"log"

	"github.com/gorm-gen/plugin/generate"
	"github.com/gorm-gen/plugin/repository"

	"example/internal/global"
	"example/internal/initialize/config"
	"example/internal/initialize/logger"
	"example/internal/initialize/mysql"
	"example/internal/models"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", global.ConfigFile, "local config file path")

	// 1、初始化本地配置文件
	config.Init(configPath)

	// 2、初始化日志
	logger.Init(false, logger.Hooks())

	// 3、初始化MySQL
	mysql.Init(false)
}

func main() {
	flag.Parse()

	opts := []generate.Option{
		generate.WithReplaceJsonTagName(map[string]generate.JsonTag{
			"deleted_at": {
				Replace: "-",
			},
		}),
	}

	g := generate.New(global.DB, opts...)

	// 1、生成models
	g.SetGenerateModel("order")

	// 2、生成gen.query
	g.SetApplyBasic(
		models.Order{},
	)

	g.Execute()

	r := repository.New(
		repository.WithModule("example"),
		repository.WithRepositoryPath("internal/repositories"),
		repository.WithGenQueryPkg("example/internal/query"),
		repository.WithGormDBVar("global.DB"),
		repository.WithGormDBVarPkg("example/internal/global"),
		repository.WithZapVar("global.Logger"),
		repository.WithZapVarPkg("example/internal/global"),
	)

	// 3、生成repository
	if err := r.Generate(
		models.Order{},
	); err != nil {
		log.Fatal(err)
		return
	}
}

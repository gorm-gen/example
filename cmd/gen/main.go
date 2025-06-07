package main

import (
	"flag"
	"log"

	"github.com/gorm-gen/plugin/generate"
	"github.com/gorm-gen/repository"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"example/internal/global"
	"example/internal/initialize/config"
	"example/internal/initialize/logger"
	"example/internal/initialize/mysql"
	"example/internal/models"
	"example/internal/query/methods"
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
	generator := g.Generator()

	// 1、生成models
	g.SetGenerateModel("order")
	g.SetGenerateModel("order_item")
	company := generator.GenerateModel("company")
	identityCard := generator.GenerateModel("identity_card")
	creditCard := generator.GenerateModel("credit_card")
	language := generator.GenerateModel("language")
	classify := generator.GenerateModel("classify")
	area := generator.GenerateModel("area")
	userLanguage := generator.GenerateModel("user_language")

	user := generator.GenerateModel("user",
		gen.FieldRelate(field.BelongsTo, "Company", company, &field.RelateConfig{
			RelatePointer: true,
			JSONTag:       "company,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"CompanyID"},
				"references": {"ID"},
			},
		}),
		gen.FieldRelate(field.HasOne, "IdentityCard", identityCard, &field.RelateConfig{
			RelatePointer: true,
			JSONTag:       "identity_card,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"UserID"},
				"references": {"ID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "CreditCards", creditCard, &field.RelateConfig{
			RelateSlicePointer: true,
			JSONTag:            "credit_cards,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"UserID"},
				"references": {"ID"},
			},
		}),
		gen.FieldRelate(field.Many2Many, "Languages", language, &field.RelateConfig{
			JSONTag: "languages,omitempty",
			GORMTag: field.GormTag{
				"many2many":      {"user_language"},
				"foreignKey":     {"ID"},
				"joinForeignKey": {"UserID"},
				"references":     {"ID"},
				"joinReferences": {"LanguageID"},
			},
		}),
	)

	languages := generator.GenerateModel("language",
		gen.FieldRelate(field.Many2Many, "Users", user, &field.RelateConfig{
			JSONTag: "users,omitempty",
			GORMTag: field.GormTag{
				"many2many":      {"user_language"},
				"foreignKey":     {"ID"},
				"joinForeignKey": {"LanguageID"},
				"references":     {"ID"},
				"joinReferences": {"UserID"},
			},
		}),
	)

	classifies := generator.GenerateModel("classify",
		gen.FieldRelate(field.BelongsTo, "Parent", classify, &field.RelateConfig{
			RelatePointer: true,
			JSONTag:       "parent,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"Pid"},
				"references": {"ID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "Child", classify, &field.RelateConfig{
			JSONTag: "child,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"Pid"},
				"references": {"ID"},
			},
		}),
	)

	areas := generator.GenerateModel("area",
		gen.FieldRelate(field.HasOne, "Parent", area, &field.RelateConfig{
			RelatePointer: true,
			JSONTag:       "parent,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"Pid"},
				"references": {"ID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "Child", area, &field.RelateConfig{
			JSONTag: "child,omitempty",
			GORMTag: field.GormTag{
				"foreignKey": {"Pid"},
				"references": {"ID"},
			},
		}),
	)

	// 2、生成gen.query
	g.SetApplyBasic(
		models.Order{},
		models.OrderItem{},
		company,
		user,
		identityCard,
		creditCard,
		languages,
		classifies,
		areas,
		userLanguage,
	)

	g.SetApplyInterface(
		func(methods.Query) {},
		user,
		company,
	)
	g.SetApplyInterface(
		func(methods.User) {},
		user,
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
		models.OrderItem{},
		models.User{},
		models.Company{},
		models.IdentityCard{},
		models.CreditCard{},
		models.Language{},
		models.Classify{},
		models.Area{},
		models.UserLanguage{},
	); err != nil {
		log.Fatal(err)
		return
	}

	// 4、生成分表repository
	if err := r.ShardingGenerate(
		"Sharding",
		models.Order{},
		models.OrderItem{},
	); err != nil {
		log.Fatal(err)
		return
	}
}

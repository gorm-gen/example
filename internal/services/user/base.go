package user

import (
	"runtime/debug"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"example/internal/global"
	"example/internal/models"
	"example/internal/query"
	"example/internal/repositories"
	"example/internal/repositories/user"
)

type User struct {
	q        *query.Query
	db       *gorm.DB
	logger   *zap.Logger
	userRepo *user.User
}

func New() *User {
	return &User{
		q:        repositories.GetQuery(),
		db:       global.DB,
		logger:   global.Logger,
		userRepo: user.New(),
	}
}

func (u *User) GetCompanyName(id int) (string, error) {
	res, err := u.q.User.GetCompanyName(id)
	if err != nil {
		if repositories.IsRealErr(err) {
			u.logger.Error("【User.GetCompanyName】失败", zap.Error(err), zap.ByteString("debug.Stack", debug.Stack()))
		}
		return "", err
	}
	return res[u.q.User.Name.ColumnName().String()].(string), nil
}

func (u *User) GetByID(id int) (models.User, error) {
	res, err := u.q.User.GetByID(id)
	if err != nil {
		if repositories.IsRealErr(err) {
			u.logger.Error("【User.GetByID】失败", zap.Error(err), zap.ByteString("debug.Stack", debug.Stack()))
		}
		return models.User{}, err
	}
	return res, nil
}

package biz

import (
	"context"

	"gendemo/dal"
	"gendemo/dal/model"
	"gendemo/dal/query"
)

var q = query.Use(dal.DB.Debug())

func Create(ctx context.Context) error {

	ud := q.User.WithContext(ctx)

	userData := &model.User{ID: 1, Name: "modi"}
	return ud.Create(userData)

}

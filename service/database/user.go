package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//CheckUser 通过UserName和password获取用户信息
func CheckUser(ctx context.Context, params *model.LoginParams) (*model.User, error) {
	rows, err := database.Table("gowatcher.user_info_table").Debug().
		Select("id, user_id, user_name, password").
		Where("user_name = ? and password = ?", params.UserName, params.Password).Rows()

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in GetUser: %v", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		logrus.Errorf("GetUser error, err: %v", err.Error())
		return nil, exceptions.ErrDBHandle
	}

	var res model.User
	for rows.Next() {
		database.ScanRows(rows, &res)
	}

	return &res, nil
}

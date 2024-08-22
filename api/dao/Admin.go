// 用户 数据层
package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/pkg/db"
)

func CheckAppNameExists(appName string) (bool, error) {
	var count int64
	err := db.Db.Model(&entity.SysAdmin{}).Where("app_name = ?", appName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateUser saves a new user to the database

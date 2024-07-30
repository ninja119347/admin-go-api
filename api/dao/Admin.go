// 用户 数据层
package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/pkg/db"
)

// 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	db.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

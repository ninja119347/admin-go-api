// 用户 数据层
package dao

import (
	"admin-go-api/api/dto"
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	"admin-go-api/pkg/db"
	"time"
)

// 用户详情
func SysAdminDetail(dto dto.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	db.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// CreateUser saves a new user to the database
func CreateUser(dto dto.CreateUserDto) error {
	user := entity.SysAdmin{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		Email:      dto.Email,
		Phone:      dto.Phone,
		Status:     1, // default status
		ID:         dto.ID,
		PostId:     dto.PostId,
		DepId:      dto.DepId,
		CreateTime: util.HTime{time.Now()},
	}
	return db.Db.Create(&user).Error
}
func SearchUser(username string) (user []entity.SysAdmin, err error) {

	err = db.Db.Where("username = ?", username).Find(&user).Error
	return user, err
}

func DeleteUser(id uint) error {
	return db.Db.Where("id = ?", id).Delete(&entity.SysAdmin{}).Error
}

func UpdateUser(dto dto.UpdateUserDto) error {
	user := entity.SysAdmin{
		ID:       dto.ID,
		Username: dto.Username,
		//Password: dto.Password,
		Email:  dto.Email,
		Phone:  dto.Phone,
		PostId: dto.PostId,
		DepId:  dto.DepId,
	}
	return db.Db.Model(&entity.SysAdmin{}).Where("id = ?", dto.ID).Updates(&user).Error

}

func SearchUserById(id uint) (user entity.SysAdmin, err error) {
	err = db.Db.Where("id = ?", id).Find(&user).Error
	return user, err
}

// 用户相关结构体
package entity

import (
	"admin-go-api/common/util"
)

// 用户模型对象
type SysAdmin struct {
	//gorm.Model
	ID         uint       `json:"id" binding:"require" gorm:"column:id;comment:'主键'" `
	PostId     int        `json:"postId" gorm:"column:post_id;comment:'岗位id';NOT NULL" `
	DepId      int        `json:"depId" gorm:"column:dept_id;comment:'部门id';NOT NULL" `
	Username   string     `json:"username" gorm:"column:username;varchar(64);comment:'用户名';NOT NULL" `
	Password   string     `json:"password" gorm:"column:password;varchar(64);comment:'密码';NOT NULL" `
	Nickname   string     `json:"nickname" gorm:"column:nickname;varchar(64);comment:'昵称'" `
	Status     int        `json:"status" gorm:"column:status;default:1;comment:'状态:1->on;2->of'" `
	Icon       string     `json:"icon" gorm:"column:icon;varchar(500);comment:'头像'"`
	Email      string     `json:"email" gorm:"column:email;varchar(100);comment:'邮箱'"`
	Phone      string     `json:"phone" gorm:"column:phone;varchar(11);comment:'手机号'"`
	Note       string     `json:"note" gorm:"column:note;varchar(500);comment:'备注'"`
	CreateTime util.HTime `json:"createTime" gorm:"column:create_time;comment:'创建时间'"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结构体
type JwtAdmin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

package dto

import "time"

// 登陆对象
type LoginDto struct {
	AppName   string `json:"app_name" validate:"required"`   //app名称
	Id        string `json:"id" validate:"required"`         //固定值223
	IdType    string `json:"id_type" validate:"required"`    //Id类型
	UltraData int64  `json:"ultra_data" validate:"required"` //附加数据
	AppId     string `json:"app_id" validate:"required"`     //AppId
}

// 发送对象
type SendDto struct {
	Users []string          `json:"users" validate:"required,dive,email"` //目标用户邮箱
	Bid   string            `json:"bid" validate:"required"`              //业务大类
	Data  map[string]string `json:"data" validate:"required"`             //bid具体实现字段
}

// 检查 UtraData 是否过期的函数
func (dto *LoginDto) IsUtraDataExpired() bool {
	currentTimestamp := time.Now().Unix()
	return currentTimestamp > dto.UltraData
}

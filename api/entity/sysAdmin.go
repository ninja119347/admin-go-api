// 用户相关结构体
package entity

// 用户模型对象
type SysAdmin struct {
	//gorm.Model
	Series   string `json:"series" gorm:"column:Series;varchar(64);comment:'系列'"`
	App_name string `json:"app_name" gorm:"column:App_name;varchar(64);comment:'app名称'"`
	Type     string `json:"type" gorm:"column:Type;varchar(64);comment:'类型'"`
	Owner    string `json:"owner" gorm:"column:Owner;varchar(64);comment:'拥有者'"`
	//CreateTime util.HTime `json:"createTime" gorm:"column:create_time;comment:'创建时间'"`
}

func (SysAdmin) TableName() string {
	return "ota"
}

// 鉴权用户结构体
type JwtAdmin struct {
	Id       string `json:"id"`
	App_name string `json:"app_name"`
	Id_type  string `json:"id_type"`
}

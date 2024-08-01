package dto

// 登陆对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image" validate:"required,min=4,max=6"` //验证码
	IdKey    string `json:"idKey" validate:"required"`             //验证码id
}

// 创建用户对象
type CreateUserDto struct {
	ID       uint   `json:"id" validate:"required"`
	PostId   int    `json:"postId" `
	DepId    int    `json:"depId" `
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
}

// 更新用户对象
type UpdateUserDto struct {
	ID       uint   `json:"id" validate:"required"`
	PostId   int    `json:"postId"`
	DepId    int    `json:"depId"`
	Username string `json:"username"`
	//TODO 可以单独做一个密码验证的方法
	//Password string `json:"password" binding:"require"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
}

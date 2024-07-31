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
	ID       uint   `json:"id" binding:"require"`
	PostId   int    `json:"postId" binding:"require"`
	DepId    int    `json:"depId" binding:"require"`
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
	Email    string `json:"email" binding:"require,email"`
	Phone    string `json:"phone" binding:"require"`
}

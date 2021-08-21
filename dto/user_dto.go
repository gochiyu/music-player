package dto

import "Essential/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:telephone`
}

//对用户的信息进行选择性显示
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

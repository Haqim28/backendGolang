package users

//
type CreateUserRequest struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Location string `json:"location" form:"location" `
	Image    string `json:"image" form:"image" `
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	IsCart   string `json:"iscart" form:"iscart" `
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" `
	Phone    string `json:"phone" form:"phone" `
	Location string `json:"location" form:"location" `
	Image    string `json:"image" form:"image" `
	Gender   string `json:"gender" form:"gender" `
	Role     string `json:"role" form:"role" `
	IsCart   int    `json:"iscart" form:"iscart"`
}

package users

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Role     string `json:"role" form:"role" validate:"role"`
	// Product  []ProductResponse `json:"product" `
}

package authdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Phone    string `gorm:"type: varchar(255)" json:"phone" validate:"required"`
	Role     string `gorm:"type: varchar(255)" json:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}

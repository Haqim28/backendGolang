package models

// Import time package here ..

// Create User struct here ..
type User struct {
	ID       int               `json:"id"`
	FullName string            `json:"fullname" gorm:"type: varchar(255)"`
	Email    string            `json:"email" gorm:"type: varchar(255)"`
	Password string            `json:"password" gorm:"type: varchar(255)"`
	Phone    string            `json:"phone" gorm:"type: varchar(255)"`
	Location string            `json:"location" gorm:"type: varchar(255)"`
	Image    string            `json:"image"`
	Gender   string            `json:"gender" gorm:"type: varchar(255)"`
	Role     string            `json:"role" gorm:"type: varchar(255)"`
	IsCart   int               `json:"iscart"`
	Product  []ProductResponse `json:"product" `
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	Location string `json:"location"`
	Email    string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

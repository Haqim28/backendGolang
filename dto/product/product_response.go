package productdto

import (
	"backend/models"
)

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
	Image string `json:"image"`
	User  string `json:"user"`
}

type ProductTransactionResponse struct {
	ID    int                         `json:"id"`
	Title string                      `json:"title"`
	Price string                      `json:"price"`
	Image string                      `json:"image"`
	User  models.UsersProfileResponse `json:"user"`
}

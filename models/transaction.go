package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID   int                  `json:"buyer_id"`
	Buyer     UsersProfileResponse `json:"buyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerID  int                  `json:"seller_id"`
	Seller    UsersProfileResponse `json:"seller" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status    string               `json:"status"`
	Subtotal  int                  `json:"subtotal"`
	CartID    int                  `json:"cart_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart      Cart                 `json:"cart"`
	CreatedAt time.Time            `json:"create_at"`
	UpdatedAt time.Time            `json:"-"`
}

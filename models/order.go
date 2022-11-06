package models

type Order struct {
	ID          int             `json:"id" gorm:"primary_key:auto_increment"`
	Qty         int             `json:"qty" form:"qty"`
	CartID      int             `json:"cart_id" form:"cart_id"`
	Cart        Cart            `json:"cart"`
	Product_id  int             `json:"product_id"`
	Product     ProductResponse `json:"product" `
	Price_order int             `json:"price_order"`
}

type OrderResponse struct {
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	NamaProduct string `json:"nama_product"`
	Qty         int    `json:"qty"`
	CartID      int    `json:"cart_id"`
	Cart        Cart   `json:"cart"`
	Price_order int    `json:"price_order"`
	// Product     []ProductResponse `json:"product" `
}

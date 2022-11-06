package cartdto

type CreateCartRequest struct {
	Qty      int `json:"qty" form:"qty" validate:"required"`
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	SubTotal int `json:"subbtotal" form:"subtotal"`
}

type UpdateCartRequest struct {
	Qty      int    `json:"qty" form:"qty"`
	UserID   int    `json:"user_id" form:"user_id"`
	SubTotal int    `json:"subtotal" form:"subtotal"`
	Status   string `json:"status" form:"status"`
}

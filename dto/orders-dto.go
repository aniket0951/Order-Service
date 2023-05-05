package dto

type CreateOrderDTO struct {
	ProductId        string `json:"prod_id" validate:"required"`
	Category         string `json:"category" validate:"required"`
	ProductSellingID string `json:"prod_selling_id" validate:"required"`
	Quantity         int64  `json:"quantity" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
	UserId           string `json:"user_id" validate:"required"`
}

type UpdateOrderStatusDTO struct {
	OrderId     string `json:"order_id" validate:"required"`
	OrderStatus string `json:"order_status" validate:"required"`
}

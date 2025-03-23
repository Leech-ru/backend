package dto

import "github.com/gogf/gf/v2/frame/g"

type CreateOrderRequest struct {
	g.Meta   `path:"/order" tags:"Order" method:"post" x-group:"Order/Create" summary:"Create a new order."`
	Customer string `json:"customer" dc:"Name of the customer" x-sort:"1"`
}

type CreateOrderResponse struct {
	Status string `json:"status" dc:"Status of the order creation" x-sort:"1"`
}

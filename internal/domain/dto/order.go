package dto

import "Leech-ru/internal/domain/types"

type CustomerInfo struct {
	FIO         string  `json:"fio" validate:"required,min=2,max=100" example:"Иванов Иван Иванович"`
	PhoneNumber string  `json:"phone_number" validate:"required" example:"+79991234567"`
	Email       string  `json:"email" validate:"required,email,min=6,max=254" example:"ivanov@example.com"`
	Address     string  `json:"address" validate:"required,min=5,max=200" example:"г. Москва, ул. Ленина, д. 1, кв. 10"`
	Comment     *string `json:"comment,omitempty" validate:"omitempty,max=500" example:"Позвоните за час до доставки"`
}

type OrderDetails struct {
	LeechSize1  *int           `json:"leech_size_1" validate:"omitempty,min=0,max=500" example:"100"`
	LeechSize2  *int           `json:"leech_size_2" validate:"omitempty,min=0,max=500" example:"200"`
	LeechSize3  *int           `json:"leech_size_3" validate:"omitempty,min=0,max=500" example:"150"`
	PackageType *types.Package `json:"package_type" validate:"required,package" example:"1"`
}

type CreateOrderRequest struct {
	CustomerInfo CustomerInfo `json:"customer_info" validate:"required"`
	OrderDetails OrderDetails `json:"order_details" validate:"required,minleechsum=50,maxleechsum=1000"`
}

type CreateOrderResponse struct {
	CustomerInfo CustomerInfo `json:"customer_info"`
	OrderDetails OrderDetails `json:"order_details"`
	TotalPrice   float64      `json:"total_price" example:"3499.50"`
}

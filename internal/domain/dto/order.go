package dto

type CustomerInfo struct {
	FIO         string  `json:"fio" validate:"required,min=2,max=100"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Address     string  `json:"address" validate:"required,min=5,max=200"`
	Comment     *string `json:"comment,omitempty" validate:"omitempty,max=500"`
}

type OrderDetails struct {
	LeechSize1  *int `json:"leech_size_1" validate:"omitempty,min=0,max=500"`
	LeechSize2  *int `json:"leech_size_2" validate:"omitempty,min=0,max=500"`
	LeechSize3  *int `json:"leech_size_3" validate:"omitempty,min=0,max=500"`
	PackageType int  `json:"package_type" validate:"required,oneof=1 2 3"`
}

type CreateOrderRequest struct {
	CustomerInfo CustomerInfo `json:"customer_info" validate:"required"`
	OrderDetails OrderDetails `json:"order_details" validate:"required,minleechsum=50,maxleechsum=1000"`
}

type CreateOrderResponse struct {
	CustomerInfo CustomerInfo `json:"customer_info"`
	OrderDetails OrderDetails `json:"order_details"`
	TotalPrice   float64      `json:"total_price"`
}

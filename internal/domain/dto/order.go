package dto

type CustomerInfo struct {
	FIO         string  `json:"fio" v:"required|length:2,100#Full name is required|Full name should be 2-100 characters long"`
	PhoneNumber string  `json:"phone_number" v:"required#Phone number is required"`
	Email       string  `json:"email" v:"required|email#Email is required|Invalid email format"`
	Address     string  `json:"address" v:"required|length:5,200#Address is required|Address should be 5-200 characters long"`
	Comment     *string `json:"comment,omitempty" v:"length:0,500#Comment should be less than 500 characters"`
}

type OrderDetails struct {
	LeechSize1  int `json:"leech_size_1" v:"required|min:0#Leech size 1 is required|Leech size 1 must be positive"`
	LeechSize2  int `json:"leech_size_2" v:"required|min:0#Leech size 2 is required|Leech size 2 must be positive"`
	LeechSize3  int `json:"leech_size_3" v:"required|min:0#Leech size 3 is required|Leech size 3 must be positive"`
	PackageType int `json:"package_type" v:"required|in:1,2,3,4#Package type is required|Invalid package type"`
}

type CreateOrderRequest struct {
	CustomerInfo CustomerInfo `json:"customer_info" v:"required#Customer information is required"`
	OrderDetails OrderDetails `v:"required|min-leech-sum:55#Order details are required|Sum of all leeches must be at least 50"`
}

type CreateOrderResponse struct {
	CustomerInfo CustomerInfo `json:"customer_info"`
	OrderDetails OrderDetails `json:"order_details"`
	TotalPrice   float64      `json:"total_price"`
}

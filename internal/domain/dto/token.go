package dto

type RefreshRequest struct {
	Goyda string `json:"goyda" form:"goyda" validate:"required,eq=true"`
}

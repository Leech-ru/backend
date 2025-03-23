package dto

type PingResponse struct {
	Message string `json:"message" dc:"Message returned in the ping response" x-sort:"1"`
}

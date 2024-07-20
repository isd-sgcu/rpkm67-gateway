package dto

type Count struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateCountRequest struct {
	Name string `json:"name"`
}

type CreateCountResponse struct {
	Count *Count `json:"count"`
}

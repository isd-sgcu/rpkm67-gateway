package dto

type Selection struct {
	Id      string   `json:"id"`
	UserId  string   `json:"user_id"`
	BaanIds []string `json:"baan_ids"`
}

type CreateSelectionRequest struct {
	UserId  string   `json:"user_id" validate:"required"`
	BaanIds []string `json:"baan_ids" validate:"required"`
}

type CreateSelectionResponse struct {
	Selection Selection `json:"selection"`
}

type FindByStudentIdSelectionRequest struct {
	UserId string `json:"user_id" validate:"required"`
}

type FindByStudentIdSelectionResponse struct {
	Selection Selection `json:"selection"`
}

type UpdateSelectionRequest struct {
	Selection Selection `json:"selection" validate:"required"`
}

type UpdateSelectionResponse struct {
	Success bool `json:"success"`
}

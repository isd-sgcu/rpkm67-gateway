package dto

type Selection struct {
	Id      string   `json:"id"`
	GroupId string   `json:"group_id"`
	BaanIds []string `json:"baan_ids"`
}

type CreateSelectionRequest struct {
	GroupId string   `json:"group_id" validate:"required"`
	BaanIds []string `json:"baan_ids" validate:"required"`
}

type CreateSelectionResponse struct {
	Selection *Selection `json:"selection"`
}

type FindByGroupIdSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required"`
}

type FindByGroupIdSelectionResponse struct {
	Selection *Selection `json:"selection"`
}

type UpdateSelectionRequest struct {
	Selection *Selection `json:"selection" validate:"required"`
}

type UpdateSelectionResponse struct {
	Success bool `json:"success"`
}

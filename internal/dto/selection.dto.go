package dto

type Selection struct {
	Id      string `json:"id"`
	GroupId string `json:"group_id"`
	BaanId  string `json:"baan_id"`
	Order   int    `json:"order"`
}

type CreateSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required"`
	BaanId  string `json:"baan_ids" validate:"required"`
	Order   int    `json:"order" validate:"required"`
}

type CreateSelectionResponse struct {
	Selection *Selection `json:"selection"`
}

type FindByGroupIdSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required"`
}

type FindByGroupIdSelectionResponse struct {
	Selections []*Selection `json:"selection"`
}

type DeleteSelectionRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeleteSelectionResponse struct {
	Success bool `json:"success"`
}

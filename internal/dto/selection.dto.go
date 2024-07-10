package dto

type Selection struct {
	GroupId string `json:"group_id"`
	BaanId  string `json:"baan_id"`
	Order   int    `json:"order"`
}

type BaanCount struct {
	BaanId string `json:"baan_id"`
	Count  int    `json:"count"`
}

type CreateSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required,uuid"`
	BaanId  string `json:"baan_id" validate:"required"`
	Order   int    `json:"order" validate:"required"`
}

type CreateSelectionResponse struct {
	Selection *Selection `json:"selection"`
}

type FindByGroupIdSelectionRequest struct {
	GroupId string `json:"group_id"`
}

type FindByGroupIdSelectionResponse struct {
	Selections []*Selection `json:"selection"`
}

type UpdateSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required,uuid"`
	BaanId  string `json:"baan_id" validate:"required"`
	Order   int    `json:"order" validate:"required"`
}

type UpdateSelectionResponse struct {
	Selection *Selection `json:"selection"`
}

type DeleteSelectionRequest struct {
	GroupId string `json:"group_id" validate:"required,uuid"`
	BaanId  string `json:"baan_id" validate:"required"`
}

type DeleteSelectionResponse struct {
	Success bool `json:"success"`
}

type CountByBaanIdSelectionResponse struct {
	BaanCounts []*BaanCount `json:"baan_counts"`
}

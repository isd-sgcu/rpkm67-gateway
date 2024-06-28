package dto

type Selection struct {
	Id      string `json:"id"`
	GroupId string `json:"group_id"`
	BaanId  string `json:"baan_id"`
	Order   int    `json:"order"`
}

type BaanCount struct {
	BaanId string `json:"baan_id"`
	Count  int32  `json:"count"`
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
	Selections []*Selection `json:"selections"`
}

type DeleteSelectionRequest struct {
	GroupId string `json:"id" validate:"required"`
}

type DeleteSelectionResponse struct {
	Success bool `json:"success"`
}

type CountByBaanIdSelectionRequest struct {
}

type CountByBaanIdSelectionResponse struct {
	BaanCounts []*BaanCount `json:"baan_counts"`
}

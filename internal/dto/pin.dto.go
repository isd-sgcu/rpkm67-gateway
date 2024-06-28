package dto

type Pin struct {
	WorkshopId string `json:"workshop_id"`
	Code       string `json:"code"`
}

type FindAllPinRequest struct {
}

type FindAllPinResponse struct {
	Pins []*Pin `json:"pins"`
}

type ResetPinRequest struct {
	WorkshopId string `json:"workshop_id" validate:"required"`
}

type ResetPinResponse struct {
	Success bool `json:"success"`
}

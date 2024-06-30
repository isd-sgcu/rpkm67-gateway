package dto

type Pin struct {
	ActivityId string `json:"activity_id"`
	Code       string `json:"code"`
}

type FindAllPinRequest struct {
}

type FindAllPinResponse struct {
	Pins []*Pin `json:"pins"`
}

type ResetPinRequest struct {
	ActivityId string `json:"activity_id"`
}

type ResetPinResponse struct {
	Success bool `json:"success"`
}

type CheckPinRequest struct {
	ActivityId string `json:"activity_id"`
	Code       string `json:"code"`
}

type CheckPinResponse struct {
	IsMatch bool `json:"is_match"`
}

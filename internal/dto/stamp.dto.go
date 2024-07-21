package dto

type Stamp struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PointA int32  `json:"point_a"`
	PointB int32  `json:"point_b"`
	PointC int32  `json:"point_c"`
	PointD int32  `json:"point_d"`
	Stamp  string `json:"stamp"`
}

type FindByUserIdStampRequest struct {
	UserID string `json:"user_id"`
}

type FindByUserIdStampResponse struct {
	Stamp *Stamp `json:"stamp"`
}

type StampByUserIdRequest struct {
	UserID     string `json:"user_id"`
	ActivityId string `json:"activity_id"`
	PinCode    string `json:"pin_code"`
	Answer     string `json:"answer"`
}

type StampByUserIdBodyRequest struct {
	ActivityId string `json:"activity_id"`
	PinCode    string `json:"pin_code"`
	Answer     string `json:"answer"`
}

type StampByUserIdResponse struct {
	Stamp *Stamp `json:"stamp"`
}

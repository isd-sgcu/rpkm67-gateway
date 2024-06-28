package dto

type CheckIn struct {
	Id     string `json:"id"`
	Event  string `json:"event"`
	Email  string `json:"email"`
	UserId string `json:"user_id"`
}

type CreateCheckInRequest struct {
	Event  string `json:"event" validate:"required"`
	Email  string `json:"email" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type CreateCheckInResponse struct {
	CheckIn *CheckIn `json:"check_in"`
}

type FindByUserIdCheckInRequest struct {
	UserId string `json:"user_id" validate:"required"`
}

type FindByUserIdCheckInResponse struct {
	CheckIns []*CheckIn `json:"check_ins"`
}

type FindByEmailCheckInRequest struct {
	Email string `json:"email" validate:"required"`
}

type FindByEmailCheckInResponse struct {
	CheckIns []*CheckIn `json:"check_ins"`
}

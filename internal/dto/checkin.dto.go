package dto

type CheckIn struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Event  string `json:"event"`
}

type CreateCheckInRequest struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Event  string `json:"event"`
}

type CreateCheckInResponse struct {
	CheckIn *CheckIn `json:"checkin"`
}

type FindByUserIdCheckInRequest struct {
	UserID string `json:"user_id"`
}

type FindByUserIdCheckInResponse struct {
	CheckIns []*CheckIn `json:"checkins"`
}

type FindByEmailCheckInRequest struct {
	Email string `json:"email"`
}

type FindByEmailCheckInResponse struct {
	CheckIns []*CheckIn `json:"checkins"`
}

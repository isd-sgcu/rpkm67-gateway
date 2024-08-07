package dto

type CheckIn struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	Event       string `json:"event"`
	Timestamp   string `json:"timestamp"`
	IsDuplicate bool   `json:"is_duplicate"`
}

type CreateCheckInRequest struct {
	UserID    string `json:"user_id"`
	StudentID string `json:"student_id"`
	Email     string `json:"email"`
	Event     string `json:"event"`
}

type CreateCheckInResponse struct {
	CheckIn   *CheckIn `json:"checkin"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
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

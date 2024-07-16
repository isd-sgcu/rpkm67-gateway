package dto

type Group struct {
	Id          string      `json:"id"`
	LeaderID    string      `json:"leader_id"`
	Token       string      `json:"token"`
	Members     []*UserInfo `json:"members"`
	IsConfirmed bool        `json:"is_confirmed"`
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageUrl  string `json:"image_url"`
}

type FindByUserIdGroupRequest struct {
	UserId string `json:"user_id"`
}

type FindByUserIdGroupResponse struct {
	Group *Group `json:"group"`
}

type FindByTokenGroupRequest struct {
	Token string `json:"token"`
}

type FindByTokenGroupResponse struct {
	Id     string    `json:"id"`
	Token  string    `json:"token"`
	Leader *UserInfo `json:"leader"`
}

type UpdateConfirmGroupBody struct {
	IsConfirmed bool `json:"is_confirmed" validate:"required"`
}

type UpdateConfirmGroupRequest struct {
	LeaderId    string `json:"leader_id"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type UpdateConfirmGroupResponse struct {
	Group *Group `json:"group"`
}

type JoinGroupRequest struct {
	Token  string `json:"token" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type JoinGroupResponse struct {
	Group *Group `json:"group"`
}

type LeaveGroupRequest struct {
	UserId string `json:"user_id" validate:"required"`
}

type LeaveGroupResponse struct {
	Group *Group `json:"group"`
}

type SwitchGroupBody struct {
	UserId        string `json:"user_id" validate:"required"`
	NewGroupToken string `json:"new_group_token" validate:"required"`
}

type SwitchGroupResponse struct {
	Group *Group `json:"group"`
}

type DeleteMemberGroupBody struct {
	RequestingUserId string `json:"requesting_user_id" validate:"required"`
	DeletedUserId    string `json:"deleted_user_id" validate:"required"`
}

type DeleteMemberGroupRequest struct {
	UserId   string `json:"user_id"`
	LeaderId string `json:"leader_id"`
}

type DeleteMemberGroupResponse struct {
	Group *Group `json:"group"`
}

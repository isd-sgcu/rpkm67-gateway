package dto

type Group struct {
	Id       string      `json:"id"`
	LeaderID string      `json:"leader_id"`
	Token    string      `json:"token"`
	Members  []*UserInfo `json:"members"`
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageUrl  string `json:"image_url"`
}

type FindOneGroupRequest struct {
	UserId string `json:"user_id"`
}

type FindOneGroupResponse struct {
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

type UpdateGroupRequest struct {
	Group    *Group `json:"group"`
	LeaderId string `json:"leader_id"`
}

type UpdateGroupResponse struct {
	Group *Group `json:"group"`
}

type JoinGroupRequest struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

type JoinGroupResponse struct {
	Group *Group `json:"group"`
}

type DeleteMemberGroupRequest struct {
	UserId   string `json:"user_id"`
	LeaderId string `json:"leader_id"`
}

type DeleteMemberGroupResponse struct {
	Group *Group `json:"group"`
}

type LeaveGroupRequest struct {
	UserId string `json:"user_id"`
}

type LeaveGroupResponse struct {
	Group *Group `json:"group"`
}

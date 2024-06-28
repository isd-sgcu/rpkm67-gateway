package dto

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-model/constant"
)

type User struct {
	Id          string        `json:"id"`
	Email       string        `json:"email"`
	Nickname    string        `json:"nickname"`
	Title       string        `json:"title"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Year        int           `json:"year"`
	Faculty     string        `json:"faculty"`
	Tel         string        `json:"tel"`
	ParentTel   string        `json:"parent_tel"`
	Parent      string        `json:"parent"`
	FoodAllergy string        `json:"food_allergy"`
	DrugAllergy string        `json:"drug_allergy"`
	Illness     string        `json:"illness"`
	Role        constant.Role `json:"role"`
	PhotoKey    string        `json:"photo_key"`
	PhotoUrl    string        `json:"photo_url"`
	Baan        string        `json:"baan"`
	ReceiveGift int           `json:"receive_gift"`
	GroupID     *uuid.UUID    `json:"group_id"`
	CheckIns    []*CheckIn    `json:"check_ins"`
}

type FindOneUserRequest struct {
	Id string `json:"id" validate:"required"`
}

type FindOneUserResponse struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
}

type UpdateUserResponse struct {
	User *User `json:"user"`
}

package user

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
)

func ProtoToDto(in *userProto.User) *dto.User {
	return &dto.User{
		Id:          in.Id,
		Email:       in.Email,
		Nickname:    in.Nickname,
		Title:       in.Title,
		Firstname:   in.Firstname,
		Lastname:    in.Lastname,
		Year:        int(in.Year),
		Faculty:     in.Faculty,
		Tel:         in.Tel,
		ParentTel:   in.ParentTel,
		Parent:      in.Parent,
		FoodAllergy: in.FoodAllergy,
		DrugAllergy: in.DrugAllergy,
		Illness:     in.Illness,
		PhotoKey:    in.PhotoKey,
		PhotoUrl:    in.PhotoUrl,
		Baan:        in.Baan,
		ReceiveGift: int(in.ReceiveGift),
		GroupId:     in.GroupId,
	}
}

func CreateUpdateUserRequestProto(req *dto.UpdateUserProfileRequest) *userProto.UpdateUserRequest {
	return &userProto.UpdateUserRequest{
		Id:          req.Id,
		Nickname:    req.Nickname,
		Title:       req.Title,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Year:        int32(req.Year),
		Faculty:     req.Faculty,
		Tel:         req.Tel,
		ParentTel:   req.ParentTel,
		Parent:      req.Parent,
		FoodAllergy: req.FoodAllergy,
		DrugAllergy: req.DrugAllergy,
		Illness:     req.Illness,
		Baan:        req.Baan,
		ReceiveGift: int32(req.ReceiveGift),
		GroupId:     req.GroupId,
	}
}

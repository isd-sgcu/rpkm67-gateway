package user

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
)

func ProtoToDto(in *userProto.User) *dto.User {
	return &dto.User{
		Id:          in.Id,
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

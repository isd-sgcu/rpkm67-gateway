package test

import (
	"github.com/bxcodec/faker/v4"
	// "github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	groupsProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/group/v1"
)

func MockUserProto() *userProto.User {
	user := &userProto.User{
		Id:          faker.UUIDDigit(),
		Email:       faker.Email(),
		Nickname:    faker.FirstName(),
		Title:       faker.Sentence(),
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		Year:        1,
		Faculty:     faker.Sentence(),
		Tel:         faker.Phonenumber(),
		ParentTel:   faker.Phonenumber(),
		Parent:      faker.Name(),
		FoodAllergy: faker.Word(),
		DrugAllergy: faker.Word(),
		Illness:     faker.Word(),
		Role:        faker.Word(),
		PhotoKey:    faker.Word(),
		PhotoUrl:    faker.URL(),
		Baan:        faker.UUIDDigit(),
		GroupId:     faker.UUIDDigit(),
		ReceiveGift: 1222,
	}
	return user
}

func MockGroupsProto() []*groupsProto.Group {
	var groups []*groupsProto.Group
	for i := 0; i < 10; i++ {
		group := &groupsProto.Group{
			Id:          faker.UUIDDigit(),
			LeaderID:    faker.UUIDDigit(),
			Token:       faker.UUIDDigit(),
			Members:     []*groupsProto.UserInfo{},
			IsConfirmed: true,
		}
		groups = append(groups, group)
	}
	return groups
}

func MockUserInfoProto() []*groupsProto.UserInfo {
	var users []*groupsProto.UserInfo
	for i := 0; i < 10; i++ {
		user := &groupsProto.UserInfo{
			Id:        faker.UUIDDigit(),
			Firstname: faker.FirstName(),
			Lastname:  faker.LastName(),
			ImageUrl:  faker.URL(),
		}
		users = append(users, user)
	}
	return users
}
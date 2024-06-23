package group

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	groupProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/group/v1"
)

func GroupProtoToDto(group *groupProto.Group) *dto.Group {
	return &dto.Group{
		Id:       group.Id,
		LeaderID: group.LeaderID,
		Token:    group.Token,
		Members:  convertMembers(group.Members),
	}
}

func GroupDtoToProto(group *dto.Group) *groupProto.Group {
	var members []*groupProto.UserInfo
	for _, member := range group.Members {
		members = append(members, &groupProto.UserInfo{
			Id:        member.Id,
			Firstname: member.FirstName,
			Lastname:  member.LastName,
			ImageUrl:  member.ImageUrl,
		})
	}
	return &groupProto.Group{
		Id:       group.Id,
		LeaderID: group.LeaderID,
		Token:    group.Token,
		Members:  members,
	}
}

func UserInfoProtoToDto(userInfo *groupProto.UserInfo) *dto.UserInfo {
	return &dto.UserInfo{
		Id:        userInfo.Id,
		FirstName: userInfo.Firstname,
		LastName:  userInfo.Lastname,
		ImageUrl:  userInfo.ImageUrl,
	}
}

func convertMembers(members []*groupProto.UserInfo) []*dto.UserInfo {
	var convertedMembers []*dto.UserInfo
	for _, member := range members {
		convertedMembers = append(convertedMembers, &dto.UserInfo{
			Id:        member.Id,
			FirstName: member.Firstname,
			LastName:  member.Lastname,
			ImageUrl:  member.ImageUrl,
		})
	}
	return convertedMembers
}

func GroupProtoToDtoList(groups []*groupProto.Group) []*dto.Group {
	var result []*dto.Group
	for _, group := range groups {
		result = append(result, GroupProtoToDto(group))
	}
	return result
}

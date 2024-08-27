package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
	"github.com/isd-sgcu/rpkm67-gateway/internal/group"
	userMock "github.com/isd-sgcu/rpkm67-gateway/mocks/user"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	objectMock "github.com/isd-sgcu/rpkm67-gateway/mocks/object"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger

	userProto *userProto.User
	userDto *dto.User

	Groups                   []*dto.Group
	Group                    *dto.Group

	FindByEmailRequestProto *userProto.FindByEmailRequest
	
	FindOneUserRequestProto *userProto.FindOneUserRequest
	FindOneUserRequestDto *dto.FindOneUserRequest
	UpdateUserProfileRequestProto *userProto.UpdateUserRequest
	UpdateUserProfileRequestDto *dto.UpdateUserProfileRequest
	UpdateUserPictureRequestDto *dto.UpdateUserPictureRequest
	ObjectService *dto.Object
	Err apperror.AppError
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	groupsProto := MockGroupsProto()
	groupProto := groupsProto[0]

	t.Groups = group.GroupProtoToDtoList(groupsProto)
	t.Group = group.GroupProtoToDto(groupProto)

	t.userProto = MockUserProto()
	t.userDto = user.ProtoToDto(t.userProto)

	t.FindOneUserRequestProto = &userProto.FindOneUserRequest{
		Id: t.userProto.Id,
	}

	t.FindOneUserRequestDto = &dto.FindOneUserRequest{
		Id: t.userDto.Id,
	}

	t.UpdateUserProfileRequestProto = &userProto.UpdateUserRequest{
		Id: t.userProto.Id,
		Email: t.userProto.Email,
		Nickname: t.userProto.Nickname,
		Title: t.userProto.Title,
		Firstname: t.userProto.Firstname,
		Lastname: t.userProto.Lastname,
		Year: t.userProto.Year,
		Faculty: t.userProto.Faculty,
		Tel: t.userProto.Tel,
		ParentTel: t.userProto.ParentTel,
		Parent: t.userProto.Parent,
		FoodAllergy: t.userProto.FoodAllergy,
		DrugAllergy: t.userProto.DrugAllergy,
		Illness: t.userProto.Illness,
		Role: t.userProto.Role,
		PhotoKey: t.userProto.PhotoKey,
		PhotoUrl: t.userProto.PhotoUrl,
		Baan: t.userProto.Baan,
		GroupId: t.userProto.GroupId,
		ReceiveGift: t.userProto.ReceiveGift,
	}

	t.UpdateUserProfileRequestDto = &dto.UpdateUserProfileRequest{
		Id: t.userDto.Id,
		Nickname: t.userDto.Nickname,
		Title: t.userDto.Title,
		Firstname: t.userDto.Firstname,
		Lastname: t.userDto.Lastname,
		Year: t.userDto.Year,
		Faculty: t.userDto.Faculty,
		Tel: t.userDto.Tel,
		ParentTel: t.userDto.ParentTel,
		Parent: t.userDto.Parent,
		FoodAllergy: t.userDto.FoodAllergy,
		DrugAllergy: t.userDto.DrugAllergy,
		Illness: t.userDto.Illness,
		PhotoKey: t.userDto.PhotoKey,
		PhotoUrl: t.userDto.PhotoUrl,
		Baan: t.userDto.Baan,
		GroupId: t.userDto.GroupId,
		ReceiveGift: t.userDto.ReceiveGift,		
	}
}

func (t *UserServiceTest) TestFindOneUserSuccess() {
	client := userMock.NewMockClient(t.controller)
	objSvc := objectMock.NewMockService(t.controller)
	svc := user.NewService(client, objSvc, t.logger)
	
	protoResp := &userProto.FindOneUserResponse{
		User: t.userProto,
	}

	createUserDto := user.ProtoToDto(protoResp.User)
	expected := &dto.FindOneUserResponse{
		User: createUserDto,
	}

	client.EXPECT().FindOne(gomock.Any(), t.FindOneUserRequestProto).Return(protoResp, nil)
	actual, err := svc.FindOne(t.FindOneUserRequestDto)

	t.Nil(err)
	t.Equal(expected, actual)
}


func (t* UserServiceTest) TestFindOneUserInvalidArgument() {
	client := userMock.NewMockClient(t.controller)
	svc := user.NewService(client, nil, t.logger)
	
	protoReq := t.FindOneUserRequestProto
	expected := apperror.BadRequest
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindOne(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindOne(t.FindOneUserRequestDto)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t* UserServiceTest) TestFindOneUserInternalError() {
	client := userMock.NewMockClient(t.controller)
	objSvc := objectMock.NewMockService(t.controller)
	svc := user.NewService(client, objSvc, t.logger)

	protoReq := t.FindOneUserRequestProto
	expected := apperror.InternalServer
	clientErr := apperror.InternalServer

	client.EXPECT().FindOne(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindOne(t.FindOneUserRequestDto)

	t.Nil(actual)
	t.Equal(expected, err)
}


func (t* UserServiceTest) TestUpdateProfileSuccess() {
	// ❌ Error
	client := userMock.NewMockClient(t.controller)
	objectSvc := objectMock.NewMockService(t.controller)
	svc := user.NewService(client, objectSvc, t.logger)

	protoResp := &userProto.UpdateUserResponse{
		User: t.userProto,
	}

	createUserDto := user.ProtoToDto(protoResp.User)
	expected := &dto.UpdateUserProfileResponse{
		User: createUserDto,
	}

	client.EXPECT().Update(gomock.Any(), t.UpdateUserProfileRequestProto).Return(protoResp, nil)
	actual, err := svc.UpdateProfile(t.UpdateUserProfileRequestDto)

	t.Nil(err)
	t.Equal(expected, actual)
}


func (t *UserServiceTest) TearDownTest() {
	t.controller.Finish()
}

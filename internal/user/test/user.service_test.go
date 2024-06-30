package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
	"github.com/isd-sgcu/rpkm67-gateway/internal/group"
	// userMock "github.com/isd-sgcu/rpkm67-gateway/mocks/user"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	// objectMock "github.com/isd-sgcu/rpkm67-gateway/mocks/object"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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
}

func (t *UserServiceTest) TestFindOneUserSuccess() {
	// client := userMock.NewMockClient(t.controller)
	// objSvc := objectMock.NewMockService(t.controller)
	// svc := user.NewService(client, objSvc, t.logger)
	
	// protoResp := &userProto.FindOneUserResponse{
	// 	User: t.userProto,
	// }

	// createUserDto := user.ProtoToDto(protoResp.User)
	// expected := &dto.FindOneUserResponse{
	// 	User: createUserDto,
	// }

	// client.EXPECT().Create(gomock.Any(), t.FindOneUserRequestProto).Return(protoResp, nil)
	// actual, err := svc.FindOne(t.FindOneUserRequestDto)

	// t.Nil(err)
	// t.Equal(expected, actual)
}

func (t *UserServiceTest) TearDownTest() {
	t.controller.Finish()
}

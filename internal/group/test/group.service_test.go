package test

import(
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/group"
	groupMock "github.com/isd-sgcu/rpkm67-gateway/mocks/group"
	groupProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/group/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type groupServiceTest struct {
	suite.Suite
	controller                     *gomock.Controller
	logger                         *zap.Logger
	groupsProto                    []*groupProto.Group
	groupsDto                      []*dto.Group
	groupProto                     *groupProto.Group
	groupDto                       *dto.Group
	Err                            apperror.AppError
}

//func DeleteMember
func TestgroupService(t *testing.T) {
	suite.Run(t, new(groupServiceTest))
	}
	
/*
func (t *groupServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.groupsProto = MockgroupsProto()
	t.groupProto = t.groupsProto[0]
	t.groupsDto = group.ProtoToDtoList(t.groupsProto)
	t.groupDto = group.ProtoToDto(t.groupProto)

	t.CreategroupProtoRequest = &groupProto.CreategroupRequest{
		GroupId: t.groupProto.GroupId,
		BaanId:  t.groupProto.BaanId,
	}
	t.CreategroupDtoRequest = &dto.CreategroupRequest{
		GroupId: t.groupDto.GroupId,
		BaanId:  t.groupDto.BaanId,
	}
	t.FindByGroupIdgroupProtoRequest = &groupProto.FindByGroupIdgroupRequest{
		GroupId: t.groupProto.GroupId,
	}
	t.FindByGroupIdgroupDtoRequest = &dto.FindByGroupIdgroupRequest{
		GroupId: t.groupDto.GroupId,
	}
	t.DeletegroupProtoRequest = &groupProto.DeletegroupRequest{
		Id: t.groupProto.Id,
	}
	t.DeletegroupDtoRequest = &dto.DeletegroupRequest{
		Id: t.groupDto.Id,
	}
}

func (t *groupServiceTest) TestCreategroupSuccess() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoResp := &groupProto.CreategroupResponse{
		group: t.groupProto,
	}

	creategroupDto := group.ProtoToDto(protoResp.group)
	expected := &dto.CreategroupResponse{
		group: creategroupDto,
	}

	client.EXPECT().Create(gomock.Any(), t.CreategroupProtoRequest).Return(protoResp, nil)
	actual, err := svc.Creategroup(t.CreategroupDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *groupServiceTest) TestCreategroupInvalidArgument() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.CreategroupProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().Create(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.Creategroup(t.CreategroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *groupServiceTest) TestCreategroupInternalError() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.CreategroupProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().Create(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.Creategroup(t.CreategroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *groupServiceTest) TestFindByGroupIdgroupSuccess() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoResp := &groupProto.FindByGroupIdgroupResponse{
		group: t.groupProto,
	}
	expected := &dto.FindByGroupIdgroupResponse{
		group: t.groupDto,
	}

	client.EXPECT().FindByGroupId(gomock.Any(), t.FindByGroupIdgroupProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByGroupIdgroup(t.FindByGroupIdgroupDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *groupServiceTest) TestFindByGroupIdgroupInvalidArgument() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.FindByGroupIdgroupProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindByGroupId(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindByGroupIdgroup(t.FindByGroupIdgroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *groupServiceTest) TestFindByGroupIdgroupInternalError() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.FindByGroupIdgroupProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().FindByGroupId(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindByGroupIdgroup(t.FindByGroupIdgroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *groupServiceTest) TestDeletegroupSuccess() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoResp := &groupProto.DeletegroupResponse{
		Success: true,
	}
	expected := &dto.DeletegroupResponse{
		Success: true,
	}

	client.EXPECT().Delete(gomock.Any(), t.DeletegroupProtoRequest).Return(protoResp, nil)
	actual, err := svc.Deletegroup(t.DeletegroupDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *groupServiceTest) TestDeletegroupInvalidArgument() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.DeletegroupProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().Delete(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.Deletegroup(t.DeletegroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *groupServiceTest) TestDeletegroupInternalError() {
	client := groupMock.NewMockClient(t.controller)
	svc := group.NewService(client, t.logger)

	protoReq := t.DeletegroupProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().Delete(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.Deletegroup(t.DeletegroupDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

//group building error with wrong code

func (t *BaanHandlerTest) TestJoinGroupSuccess() {
	baanSvc := groupmock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := baan.NewHandler(baanSvc, validator, t.logger)

	expectedResp := &dto.FindOneBaanResponse{
		Baan: t.Baan,
	}

	context.EXPECT().Param("id").Return(t.ParamMock)
	validator.EXPECT().Validate(t.FindOneBaanReq).Return(nil)
	baanSvc.EXPECT().FindOneBaan(t.FindOneBaanReq).Return(expectedResp, t.Err)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindOneBaan(context)
}
*/
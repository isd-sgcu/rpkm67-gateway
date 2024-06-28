package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/selection"
	selectionMock "github.com/isd-sgcu/rpkm67-gateway/mocks/selection"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SelectionServiceTest struct {
	suite.Suite
	controller                         *gomock.Controller
	logger                             *zap.Logger
	SelectionsProto                    []*selectionProto.Selection
	SelectionsDto                      []*dto.Selection
	SelectionProto                     *selectionProto.Selection
	SelectionDto                       *dto.Selection
	CreateSelectionProtoRequest        *selectionProto.CreateSelectionRequest
	CreateSelectionDtoRequest          *dto.CreateSelectionRequest
	FindByGroupIdSelectionProtoRequest *selectionProto.FindByGroupIdSelectionRequest
	FindByGroupIdSelectionDtoRequest   *dto.FindByGroupIdSelectionRequest
	DeleteSelectionProtoRequest        *selectionProto.DeleteSelectionRequest
	DeleteSelectionDtoRequest          *dto.DeleteSelectionRequest
}

func TestSelectionService(t *testing.T) {
	suite.Run(t, new(SelectionServiceTest))
}

func (t *SelectionServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.SelectionsProto = MockSelectionsProto()
	t.SelectionProto = t.SelectionsProto[0]
	t.SelectionsDto = selection.ProtoToDtos(t.SelectionsProto)
	t.SelectionDto = selection.ProtoToDto(t.SelectionProto)

	t.CreateSelectionProtoRequest = &selectionProto.CreateSelectionRequest{
		GroupId: t.SelectionProto.GroupId,
		BaanId:  t.SelectionProto.BaanId,
	}
	t.CreateSelectionDtoRequest = &dto.CreateSelectionRequest{
		GroupId: t.SelectionDto.GroupId,
		BaanId:  t.SelectionDto.BaanId,
	}
	t.FindByGroupIdSelectionProtoRequest = &selectionProto.FindByGroupIdSelectionRequest{
		GroupId: t.SelectionProto.GroupId,
	}
	t.FindByGroupIdSelectionDtoRequest = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.SelectionDto.GroupId,
	}
	t.DeleteSelectionProtoRequest = &selectionProto.DeleteSelectionRequest{
		GroupId: t.SelectionProto.GroupId,
	}
	t.DeleteSelectionDtoRequest = &dto.DeleteSelectionRequest{
		GroupId: t.SelectionDto.GroupId,
	}
}

func (t *SelectionServiceTest) TestCreateSelectionSuccess() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoResp := &selectionProto.CreateSelectionResponse{
		Selection: t.SelectionProto,
	}

	createSelectionDto := selection.ProtoToDto(protoResp.Selection)
	expected := &dto.CreateSelectionResponse{
		Selection: createSelectionDto,
	}

	client.EXPECT().Create(gomock.Any(), t.CreateSelectionProtoRequest).Return(protoResp, nil)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestCreateSelectionInvalidArgument() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.CreateSelectionProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().Create(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestCreateSelectionInternalError() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.CreateSelectionProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().Create(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionSuccess() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoResp := &selectionProto.FindByGroupIdSelectionResponse{
		Selections: t.SelectionsProto,
	}
	expected := &dto.FindByGroupIdSelectionResponse{
		Selections: t.SelectionsDto,
	}

	client.EXPECT().FindByGroupId(gomock.Any(), t.FindByGroupIdSelectionProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionInvalidArgument() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.FindByGroupIdSelectionProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindByGroupId(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionInternalError() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.FindByGroupIdSelectionProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().FindByGroupId(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestDeleteSelectionSuccess() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoResp := &selectionProto.DeleteSelectionResponse{
		Success: true,
	}
	expected := &dto.DeleteSelectionResponse{
		Success: true,
	}

	client.EXPECT().Delete(gomock.Any(), t.DeleteSelectionProtoRequest).Return(protoResp, nil)
	actual, err := svc.DeleteSelection(t.DeleteSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestDeleteSelectionInvalidArgument() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.DeleteSelectionProtoRequest
	expected := apperror.BadRequestError("Invalid argument")
	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().Delete(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.DeleteSelection(t.DeleteSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestDeleteSelectionInternalError() {
	client := selectionMock.NewMockClient(t.controller)
	svc := selection.NewService(client, t.logger)

	protoReq := t.DeleteSelectionProtoRequest
	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")
	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().Delete(gomock.Any(), protoReq).Return(nil, clientErr)
	actual, err := svc.DeleteSelection(t.DeleteSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

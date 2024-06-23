package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/selection"
	selectionMock "github.com/isd-sgcu/rpkm67-gateway/mocks/client/selection"
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
	UpdateSelectionProtoRequest        *selectionProto.UpdateSelectionRequest
	UpdateSelectionDtoRequest          *dto.UpdateSelectionRequest
	Err                                apperror.AppError
}

func TestSelectionService(t *testing.T) {
	suite.Run(t, new(SelectionServiceTest))
}

func (t *SelectionServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.SelectionsProto = MockSelectionsProto()
	t.SelectionProto = t.SelectionsProto[0]
	t.SelectionsDto = selection.ProtoToDtoList(t.SelectionsProto)
	t.SelectionDto = selection.ProtoToDto(t.SelectionProto)

	t.CreateSelectionProtoRequest = &selectionProto.CreateSelectionRequest{
		GroupId: t.SelectionProto.GroupId,
		BaanIds: t.SelectionProto.BaanIds,
	}
	t.CreateSelectionDtoRequest = &dto.CreateSelectionRequest{
		GroupId: t.SelectionDto.GroupId,
		BaanIds: t.SelectionDto.BaanIds,
	}
	t.FindByGroupIdSelectionProtoRequest = &selectionProto.FindByGroupIdSelectionRequest{
		GroupId: t.SelectionProto.GroupId,
	}
	t.FindByGroupIdSelectionDtoRequest = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.SelectionDto.GroupId,
	}
	t.UpdateSelectionProtoRequest = &selectionProto.UpdateSelectionRequest{
		Selection: t.SelectionProto,
	}
	t.UpdateSelectionDtoRequest = &dto.UpdateSelectionRequest{
		Selection: t.SelectionDto,
	}
}

func (t *SelectionServiceTest) TestCreateSelectionSuccess() {
	protoResp := &selectionProto.CreateSelectionResponse{
		Selection: t.SelectionProto,
	}

	createSelectionDto := selection.ProtoToDto(protoResp.Selection)

	expected := &dto.CreateSelectionResponse{
		Selection: createSelectionDto,
	}

	client := selectionMock.SelectionClientMock{}
	client.On("Create", t.CreateSelectionProtoRequest).Return(protoResp, nil)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestCreateSelectionInvalidArgument() {
	protoReq := t.CreateSelectionProtoRequest

	expected := apperror.BadRequestError("Invalid argument")

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("Create", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestCreateSelectionInternalError() {
	protoReq := t.CreateSelectionProtoRequest

	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("Create", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.CreateSelection(t.CreateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionSuccess() {
	protoResp := &selectionProto.FindByGroupIdSelectionResponse{
		Selection: t.SelectionProto,
	}

	expected := &dto.FindByGroupIdSelectionResponse{
		Selection: t.SelectionDto,
	}

	client := selectionMock.SelectionClientMock{}
	client.On("FindByGroupId", t.FindByGroupIdSelectionProtoRequest).Return(protoResp, nil)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionInvalidArgument() {
	protoReq := t.FindByGroupIdSelectionProtoRequest

	expected := apperror.BadRequestError("Invalid argument")

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("FindByGroupId", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestFindByGroupIdSelectionInternalError() {
	protoReq := t.FindByGroupIdSelectionProtoRequest

	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("FindByGroupId", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.FindByGroupIdSelection(t.FindByGroupIdSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)

}

func (t *SelectionServiceTest) TestUpdateSelectionSuccess() {
	protoResp := &selectionProto.UpdateSelectionResponse{
		Success: true,
	}

	expected := &dto.UpdateSelectionResponse{
		Success: true,
	}

	client := selectionMock.SelectionClientMock{}
	client.On("Update", t.UpdateSelectionProtoRequest).Return(protoResp, nil)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.UpdateSelection(t.UpdateSelectionDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *SelectionServiceTest) TestUpdateSelectionInvalidArgument() {
	protoReq := t.UpdateSelectionProtoRequest

	expected := apperror.BadRequestError("Invalid argument")

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("Update", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.UpdateSelection(t.UpdateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

func (t *SelectionServiceTest) TestUpdateSelectionInternalError() {
	protoReq := t.UpdateSelectionProtoRequest

	expected := apperror.InternalServerError("rpc error: code = Internal desc = Internal error")

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client := selectionMock.SelectionClientMock{}
	client.On("Update", protoReq).Return(nil, clientErr)

	svc := selection.NewService(&client, t.logger)
	actual, err := svc.UpdateSelection(t.UpdateSelectionDtoRequest)

	t.Nil(actual)
	t.Equal(expected, err)
}

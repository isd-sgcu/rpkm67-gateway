package test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/stamp"
	pinMock "github.com/isd-sgcu/rpkm67-gateway/mocks/pin"
	stampMock "github.com/isd-sgcu/rpkm67-gateway/mocks/stamp"
	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StampServiceTest struct {
	suite.Suite
	controller                *gomock.Controller
	logger                    *zap.Logger
	stampProto                *stampProto.Stamp
	stampDto                  *dto.Stamp
	findByUserIdProtoRequest  *stampProto.FindByUserIdStampRequest
	findByUserIdDtoRequest    *dto.FindByUserIdStampRequest
	stampByUserIdProtoRequest *stampProto.StampByUserIdRequest
	stampByUserIdDtoRequest   *dto.StampByUserIdRequest
}

func TestStampService(t *testing.T) {
	suite.Run(t, new(StampServiceTest))
}

func (t *StampServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	pin := faker.Word()

	t.stampProto = &stampProto.Stamp{
		Id:     faker.UUIDDigit(),
		UserId: faker.UUIDDigit(),
		PointA: 1,
		PointB: 2,
		PointC: 3,
		PointD: 4,
		Stamp:  faker.Word(),
	}
	t.stampDto = stamp.ProtoToDto(t.stampProto)
	t.findByUserIdProtoRequest = &stampProto.FindByUserIdStampRequest{
		UserId: t.stampProto.UserId,
	}
	t.findByUserIdDtoRequest = &dto.FindByUserIdStampRequest{
		UserID: t.stampDto.UserID,
	}
	t.stampByUserIdProtoRequest = &stampProto.StampByUserIdRequest{
		UserId:     t.stampProto.UserId,
		ActivityId: pin,
	}
	t.stampByUserIdDtoRequest = &dto.StampByUserIdRequest{
		UserID:     t.stampDto.UserID,
		ActivityId: t.stampByUserIdProtoRequest.ActivityId,
		Pin:        pin,
	}
}

func (t *StampServiceTest) TestFindByUserIdSuccess() {
	client := stampMock.NewMockClient(t.controller)
	svc := stamp.NewService(client, nil, t.logger)

	protoResp := &stampProto.FindByUserIdStampResponse{
		Stamp: t.stampProto,
	}

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByUserId(t.findByUserIdDtoRequest)

	t.Nil(err)
	t.Equal(t.stampDto, actual.Stamp)
}

func (t *StampServiceTest) TestFindByUserIdStampInvalidArgument() {
	client := stampMock.NewMockClient(t.controller)
	svc := stamp.NewService(client, nil, t.logger)

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdProtoRequest).Return(nil, clientErr)

	actual, err := svc.FindByUserId(t.findByUserIdDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *StampServiceTest) TestFindByUserIdStampInternalError() {
	client := stampMock.NewMockClient(t.controller)
	svc := stamp.NewService(client, nil, t.logger)

	clientErr := apperror.InternalServer

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdProtoRequest).Return(nil, clientErr)

	actual, err := svc.FindByUserId(t.findByUserIdDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

func (t *StampServiceTest) TestStampByUserIdSuccess() {
	client := stampMock.NewMockClient(t.controller)
	pinSvc := pinMock.NewMockService(t.controller)
	svc := stamp.NewService(client, pinSvc, t.logger)

	pins := &dto.FindAllPinResponse{
		Pins: []*dto.Pin{
			{
				Code: t.stampByUserIdDtoRequest.Pin,
			},
		},
	}
	protoResp := &stampProto.StampByUserIdResponse{
		Stamp: t.stampProto,
	}

	client.EXPECT().StampByUserId(gomock.Any(), t.stampByUserIdProtoRequest).Return(protoResp, nil)
	pinSvc.EXPECT().FindAll(gomock.Any()).Return(pins, nil)

	actual, err := svc.StampByUserId(t.stampByUserIdDtoRequest)

	t.Nil(err)
	t.Equal(t.stampDto, actual.Stamp)
}

func (t *StampServiceTest) TestStampByUserIdPinNotFound() {
	client := stampMock.NewMockClient(t.controller)
	pinSvc := pinMock.NewMockService(t.controller)
	svc := stamp.NewService(client, pinSvc, t.logger)

	pins := &dto.FindAllPinResponse{
		Pins: []*dto.Pin{},
	}

	pinSvc.EXPECT().FindAll(gomock.Any()).Return(pins, nil)

	actual, err := svc.StampByUserId(t.stampByUserIdDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *StampServiceTest) TestStampByUserIdInvalidArgument() {
	client := stampMock.NewMockClient(t.controller)
	pinSvc := pinMock.NewMockService(t.controller)
	svc := stamp.NewService(client, pinSvc, t.logger)

	pins := &dto.FindAllPinResponse{
		Pins: []*dto.Pin{
			{
				Code: t.stampByUserIdDtoRequest.Pin,
			},
		},
	}

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	pinSvc.EXPECT().FindAll(gomock.Any()).Return(pins, nil)
	client.EXPECT().StampByUserId(gomock.Any(), t.stampByUserIdProtoRequest).Return(nil, clientErr)

	actual, err := svc.StampByUserId(t.stampByUserIdDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *StampServiceTest) TestStampByUserIdStampInternalError() {
	client := stampMock.NewMockClient(t.controller)
	pinSvc := pinMock.NewMockService(t.controller)
	svc := stamp.NewService(client, pinSvc, t.logger)

	pins := &dto.FindAllPinResponse{
		Pins: []*dto.Pin{
			{
				Code: t.stampByUserIdDtoRequest.Pin,
			},
		},
	}
	clientErr := apperror.InternalServer

	pinSvc.EXPECT().FindAll(gomock.Any()).Return(pins, nil)
	client.EXPECT().StampByUserId(gomock.Any(), t.stampByUserIdProtoRequest).Return(nil, clientErr)

	actual, err := svc.StampByUserId(t.stampByUserIdDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

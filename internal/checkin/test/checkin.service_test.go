package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/config"
	"github.com/isd-sgcu/rpkm67-gateway/internal/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	checkinMock "github.com/isd-sgcu/rpkm67-gateway/mocks/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/tracer"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckInServiceTest struct {
	suite.Suite
	controller                      *gomock.Controller
	logger                          *zap.Logger
	tracer                          trace.Tracer
	checkinsProto                   []*checkinProto.CheckIn
	checkinProto                    *checkinProto.CheckIn
	checkinsDto                     []*dto.CheckIn
	checkinDto                      *dto.CheckIn
	createCheckInProtoRequest       *checkinProto.CreateCheckInRequest
	createCheckInDtoRequest         *dto.CreateCheckInRequest
	findByUserIdCheckInProtoRequest *checkinProto.FindByUserIdCheckInRequest
	findByUserIdCheckInDtoRequest   *dto.FindByUserIdCheckInRequest
	findByEmailCheckInProtoRequest  *checkinProto.FindByEmailCheckInRequest
	findByEmailCheckInDtoRequest    *dto.FindByEmailCheckInRequest
}

func TestCheckInService(t *testing.T) {
	suite.Run(t, new(CheckInServiceTest))
}

func (t *CheckInServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	tracer, err := tracer.New(&config.Config{})
	if err != nil {
		t.T().Fatal(err)
	}
	t.tracer = tracer.Tracer("test")

	t.checkinsProto = MockCheckInsProto()
	t.checkinProto = t.checkinsProto[0]
	t.checkinsDto = checkin.ProtoToDtos(t.checkinsProto)
	t.checkinDto = t.checkinsDto[0]
	t.createCheckInProtoRequest = &checkinProto.CreateCheckInRequest{
		Email:  t.checkinProto.Email,
		UserId: t.checkinProto.UserId,
		Event:  t.checkinProto.Event,
	}
	t.createCheckInDtoRequest = &dto.CreateCheckInRequest{
		Email:  t.checkinDto.Email,
		UserID: t.checkinDto.UserID,
		Event:  t.checkinDto.Event,
	}
	t.findByUserIdCheckInProtoRequest = &checkinProto.FindByUserIdCheckInRequest{
		UserId: t.checkinProto.UserId,
	}
	t.findByUserIdCheckInDtoRequest = &dto.FindByUserIdCheckInRequest{
		UserID: t.checkinDto.UserID,
	}
	t.findByEmailCheckInProtoRequest = &checkinProto.FindByEmailCheckInRequest{
		Email: t.checkinProto.Email,
	}
	t.findByEmailCheckInDtoRequest = &dto.FindByEmailCheckInRequest{
		Email: t.checkinDto.Email,
	}
}

func (t *CheckInServiceTest) TestCreateCheckInSuccess() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	protoResp := &checkinProto.CreateCheckInResponse{
		CheckIn: t.checkinProto,
	}

	createCheckinDto := checkin.ProtoToDto(t.checkinProto)
	expected := &dto.CreateCheckInResponse{
		CheckIn: createCheckinDto,
	}

	client.EXPECT().Create(gomock.Any(), t.createCheckInProtoRequest).Return(protoResp, nil)
	actual, err := svc.Create(context.Background(), t.createCheckInDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *CheckInServiceTest) TestCreateCheckInInvalidArgument() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().Create(gomock.Any(), t.createCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.Create(context.Background(), t.createCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *CheckInServiceTest) TestCreateCheckInInternalError() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().Create(gomock.Any(), t.createCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.Create(context.Background(), t.createCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

func (t *CheckInServiceTest) TestFindByUserIdCheckInSuccess() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	protoResp := &checkinProto.FindByUserIdCheckInResponse{
		CheckIns: t.checkinsProto,
	}
	expected := &dto.FindByUserIdCheckInResponse{
		CheckIns: t.checkinsDto,
	}

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdCheckInProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByUserID(context.Background(), t.findByUserIdCheckInDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *CheckInServiceTest) TestFindByUserIdCheckInInvalidArgument() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByUserID(context.Background(), t.findByUserIdCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *CheckInServiceTest) TestFindByUserIdCheckInInternalError() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().FindByUserId(gomock.Any(), t.findByUserIdCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByUserID(context.Background(), t.findByUserIdCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

func (t *CheckInServiceTest) TestFindByEmailCheckInSuccess() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	protoResp := &checkinProto.FindByEmailCheckInResponse{
		CheckIns: t.checkinsProto,
	}
	expected := &dto.FindByEmailCheckInResponse{
		CheckIns: t.checkinsDto,
	}

	client.EXPECT().FindByEmail(gomock.Any(), t.findByEmailCheckInProtoRequest).Return(protoResp, nil)
	actual, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInDtoRequest)

	t.Nil(err)
	t.Equal(expected, actual)
}

func (t *CheckInServiceTest) TestFindByEmailCheckInInvalidArgument() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.InvalidArgument, apperror.BadRequest.Error())

	client.EXPECT().FindByEmail(gomock.Any(), t.findByEmailCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.BadRequest, err)
}

func (t *CheckInServiceTest) TestFindByEmailCheckInInternalError() {
	client := checkinMock.NewMockClient(t.controller)
	svc := checkin.NewService(client, t.logger, t.tracer)

	clientErr := status.Error(codes.Internal, apperror.InternalServer.Error())

	client.EXPECT().FindByEmail(gomock.Any(), t.findByEmailCheckInProtoRequest).Return(nil, clientErr)
	actual, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInDtoRequest)

	t.Nil(actual)
	t.Equal(apperror.InternalServer, err)
}

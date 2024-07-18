package test
import (
	"testing"
	ctxMock "github.com/isd-sgcu/rpkm67-gateway/mocks/context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	mockPin "github.com/isd-sgcu/rpkm67-gateway/mocks/pin"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"go.uber.org/zap"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"

)


type PinHandlerTest struct {
 	suite.Suite
 	controller         *gomock.Controller
 	logger             *zap.Logger
 	pinsProto          []*pinProto.Pin
 	pinProto           *pinProto.Pin
 	pinsDto            []*dto.Pin
 	pinDto             *dto.Pin
 	FindAllPinProtoReq *pinProto.FindAllPinRequest
 	FindAllPinDtoReq   *dto.FindAllPinRequest
	ResetPinProtoReq    *pinProto.ResetPinRequest
	ResetPinDtoReq    *dto.ResetPinRequest
	Err                apperror.AppError
	CheckPinProtoReq  *pinProto.CheckPinRequest
	CheckPinDtoReq    *dto.CheckPinRequest
 }

func TestPinHandler(t *testing.T) {
	suite.Run(t, new(PinServiceTest))
}

func (t *PinHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
 	t.logger = zap.NewNop()
	t.pinsProto = MockPinsProto()
	t.pinProto = (t.pinsProto)[0]
	t.pinsDto = pin.ProtoToDtoList(t.pinsProto)
	t.pinDto = (t.pinsDto)[0]
	t.FindAllPinProtoReq = &pinProto.FindAllPinRequest{}
 	t.FindAllPinDtoReq = &dto.FindAllPinRequest{}
	t.ResetPinProtoReq = &pinProto.ResetPinRequest{
		ActivityId: t.pinProto.ActivityId,
	}
	t.ResetPinDtoReq = &dto.ResetPinRequest{
		ActivityId: t.pinDto.ActivityId,
	}
	t.CheckPinDtoReq = &dto.CheckPinRequest{
		ActivityId: t.pinDto.ActivityId,
		Code: t.pinDto.Code,
	}
	t.CheckPinProtoReq = &pinProto.CheckPinRequest{

	}
}

func (t *PinHandlerTest) TestFindAllSuccess() {
	pinSvc := mockPin.NewMockPinServiceClient(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := pin.NewHandler(pinSvc, validator, t.logger)

	expectedResp := &dto.CreateSelectionResponse{
		Selection: t.pin,
	}

	context.EXPECT().GetString("userId").Return(t.userId)
	groupSvc.EXPECT().FindByUserId(&dto.FindByUserIdGroupRequest{UserId: t.userId}).
		Return(&dto.FindByUserIdGroupResponse{Group: &dto.Group{LeaderID: t.userId}}, nil)
	context.EXPECT().Next()
	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).SetArg(0, *t.CreateSelectionReq)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().Create(t.CreateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.Create(context)

	client := mockPin.NewMockPinServiceClient(t.controller)
	svc := pin.NewService(client, t.logger)
	
	protoResp := &pinProto.FindAllPinResponse{
		Pins: t.pinsProto,
	}
	expected := &dto.FindAllPinResponse{
		Pins: t.pinsDto,
	}
	
	client.EXPECT().FindAll(gomock.Any(), t.FindAllPinProtoReq).Return(protoResp, nil)
	
	actual, err := svc.FindAll(t.FindAllPinDtoReq)
	
	t.Nil(err)
	t.Equal(expected, actual)
}
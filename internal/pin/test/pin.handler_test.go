package test
import (
	//"context"
	"testing"

	//"time"

	"github.com/golang/mock/gomock"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	mockPin "github.com/isd-sgcu/rpkm67-gateway/mocks/pin"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type PinServiceTest struct {
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

func TestPinService(t *testing.T) {
	suite.Run(t, new(PinServiceTest))
}

func (t *PinServiceTest) SetupTest() {
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

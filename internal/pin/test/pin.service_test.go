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
	ResetPinProtoReq     *dto.ResetPinRequest
	ResetPinDtoReq    *dto.ResetPinRequest
 	Err                apperror.AppError
 }

func TestPinService(t *testing.T) {
	suite.Run(t, new(PinServiceTest))
}

func (t *PinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
 	t.logger = zap.NewNop()
	t.pinsProto = MockPinsProto()
	t.pinProto = t.pinsProto[0]
	t.pinsDto = pin.ProtoToDtoList(t.pinsProto)
	t.pinDto = t.pinsDto[0]
	t.FindAllPinProtoReq = &pinProto.FindAllPinRequest{}
 	t.FindAllPinDtoReq = &dto.FindAllPinRequest{}
	t.ResetPinProtoReq = &pinProto.ResetPinRequest{
		ActivityId: "",
	}
	t.ResetPinDtoReq = &dto.ResetPinRequest{
		ActivityId: t.pinDto.ActivityId,
	}
 	t.Err = t.Err
}


func (t *PinServiceTest) TestFindAllSuccess() {
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
	
	/*protoResp := &pinProto.FindAllPinResponse{
			Pins: []*pinProto.Pin{
				{ActivityId: "1", Code: "000000"},
				{ActivityId: "2", Code: "111111"},
				{ActivityId: "3", Code: "222222"},
				{ActivityId: "4", Code: "333333"},
				{ActivityId: "5", Code: "444444"},
				{ActivityId: "6", Code: "555555"},
		}}
	*/
    //t.NotNil(res)
    //t.Len(res.Pins, 6)
    //t.Equal("1", res.Pins[0].ActivityId)
    //t.Equal("000000", res.Pins[0].Code)
}



func (t *PinServiceTest) TestResetPin(req *dto.ResetPinRequest) {
	client := mockPin.NewMockPinServiceClient(t.controller)
	svc := pin.NewService(client, t.logger)
	
	protoResp := &pinProto.ResetPinResponse{
		Pin: t.pinProto,
	}
	expected := &dto.ResetPinResponse{
		Pin: t.pinDto,
	}
	
	client.EXPECT().FindAll(gomock.Any(), t.FindAllPinProtoReq).Return(protoResp, nil)
	
	actual, err := svc.ResetPin(t.ResetPinDtoReq)

	t.Nil(err)
	t.Equal(expected, actual)
}

/*

func TestCheckPin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockPin.NewMockPinServiceClient(ctrl)
	logger := zap.NewExample()
	service := pin.NewService(mockClient, logger)

	mockClient.EXPECT().
		CheckPin(gomock.Any(), gomock.Any()).
		Return(&pinProto.CheckPinResponse{
			IsMatch: true,
		}, nil)

	req := &dto.CheckPinRequest{ActivityId: "activity123", Code: "code123"}
	res, err := service.CheckPin(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.IsMatch)
}


/*package test

import (
	//"context"
	"testing"
	//"time"

	"github.com/golang/mock/gomock"
	//"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	mockPinProto "github.com/isd-sgcu/rpkm67-gateway/mocks/pin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type PinServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger

}

func (t *PinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *PinServiceTest) TearDownTest() {
	t.controller.Finish()
}

func (t *PinServiceTest) TestResetPinSuccess() {
	mockClient := mockPinProto.NewMockPinServiceClient(t.controller)
	service := pin.NewService(mockClient, t.logger)

	req := &dto.ResetPinRequest{
		activity_id: "123123",
	}

	res := &pinProto.ResetPinResponse{
		Success: true,
	}

	mockClient.EXPECT().ResetPin(gomock.Any(), gomock.Any()).Return(res, nil).Times(1)

	resp, err := service.ResetPin(req)

	t.NoError(err)
	t.NotNil(resp)
	t.Equal(true, resp.Success)
}

func TestPinServiceTest(t *testing.T) {
	suite.Run(t, new(PinServiceTest))
}


/*package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/pin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	pinMock "github.com/isd-sgcu/rpkm67-gateway/mocks/pin"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type PinServiceTest struct {
 	suite.Suite
 	controller          *gomock.Controller
 	logger              *zap.Logger
 	pinsProto          []*pinProto.Pin
 	pinProto           *pinProto.Pin
 	pinsDto            []*dto.Pin
 	pinDto             *dto.Pin
 	FindAllPinProtoReq *pinProto.FindAllPinRequest
 	FindAllPinDtoReq   *dto.FindAllPinRequest
	// 	FindOnePinProtoReq *pinProto.FindOnePinRequest
// 	FindOnePinDtoReq   *dto.FindOnePinRequest
 	Err                 apperror.AppError
 }

func TestpinService(t *testing.T) {
	suite.Run(t, new(PinServiceTest))
}

func (t *PinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
 	t.logger = zap.NewNop()

// 	t.pinsProto = MockPinsProto()
// 	t.pinProto = t.pinsProto[0]
 	t.pinsDto = pin.ProtoToDtoList(t.pinsProto)
 	t.pinDto = pin.ProtoToDto(t.pinProto)

 	t.FindAllPinProtoReq = &pinProto.FindAllPinRequest{}
// 	t.FindOnepinProtoReq = &pinProto.FindAllPinRequest{
// 		Id: t.pinProto.Id,
// 	}
 	t.FindAllPinDtoReq = &dto.FindAllPinRequest{}
// 	t.FindOnePinDtoReq = &dto.FindOnePinRequest{
// 		Id: t.pinDto.Id,
// }
}


func (t *PinServiceTest) TestFindAllSuccess() {
	
	mockClient := mockPinProto.NewMockPinServiceClient(t.controller)
	svc := pin.NewService(&client, t.logger)
	actual, err := svc.FindAll(t.FindAllPinDtoReq)
	expected := &dto.FindAllPinResponse{}
	t.Nil(err)
	t.Equal(expected, actual)

	// client := pinMock.PinClientMock{}
	//protoResp := &pinProto.FindAllPinResponse{
	//	pins: t.PinsProto,
	//}
 	//svc := pin.NewService(&client, t.logger)

 	//findAllpinsDto := pin.ProtoToDtoList(protoResp.pins)
 	//expected := &dto.FindAllPinResponse{
 	//	pins: findAllpinsDto,
 	//}

 	//client.On("FindAllpin", t.FindAllPinProtoReq).Return(protoResp, nil)
 	//actual, err := svc.FindAllPin(t.FindAllPinDtoReq)

 	//t.Nil(err)
 	//t.Equal(expected, actual)
 }


/*
func (t *PinServiceTest) TestResetPinSuccess() {
	//client := pinMock.NewMockService(t.controller)
	client := pinMock.NewMockService(t.controller)
	
	
	actual, err := svc.FindAll(t.FindAllPinDtoReq)
	expected := &dto.FindAllPinResponse{}
	t.Nil(err)
	t.Equal(expected, actual)

	client := pinMock.pinClientMock{}
	svc := pin.NewService(&client, t.logger)

	protoResp := &pinProto.ResetpinResponse{
		Success: true,
	}
	expected := &dto.ResetpinResponse{
		Success: true,
	}

	client.On("Resetpin", t.ResetpinProtoReq).Return(protoResp, nil)
	actual, err := svc.Resetpin(t.ResetpinDtoReq)

	t.Nil(err)
	t.Equal(expected, actual)

}
*/
// func (t *pinServiceTest) TestFindOnepinSuccess() {
// 	client := pinMock.pinClientMock{}
// 	svc := pin.NewService(&client, t.logger)

// 	protoResp := &pinProto.FindOnepinResponse{
// 		pin: t.pinProto,
// 	}
// 	expected := &dto.FindOnepinResponse{
// 		pin: t.pinDto,
// 	}

// 	client.On("FindOnepin", t.FindOnepinProtoReq).Return(protoResp, nil)
// 	actual, err := svc.FindOnepin(t.FindOnepinDtoReq)

// 	t.Nil(err)
// 	t.Equal(expected, actual)
// }

// func (t *pinServiceTest) TearDownTest() {
// 	t.controller.Finish()
// }

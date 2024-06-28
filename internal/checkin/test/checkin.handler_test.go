package test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	checkinMock "github.com/isd-sgcu/rpkm67-gateway/mocks/checkin"
	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
)

type CheckInHandlerTest struct {
	suite.Suite
	controller             *gomock.Controller
	logger                 *zap.Logger
	checkins               []*dto.CheckIn
	checkin                *dto.CheckIn
	createCheckinReq       *dto.CreateCheckInRequest
	findByUserIdCheckinReq *dto.FindByUserIdCheckInRequest
	findByEmailCheckinReq  *dto.FindByEmailCheckInRequest
}

func TestCheckinHandler(t *testing.T) {
	suite.Run(t, new(CheckInHandlerTest))
}

func (t *CheckInHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	checkinsProto := MockCheckInsProto()
	t.checkins = checkin.ProtoToDtos(checkinsProto)
	t.checkin = t.checkins[0]

	t.createCheckinReq = &dto.CreateCheckInRequest{
		Email:  t.checkin.Email,
		UserID: t.checkin.UserID,
		Event:  t.checkin.Event,
	}
	t.findByUserIdCheckinReq = &dto.FindByUserIdCheckInRequest{
		UserID: t.checkin.UserID,
	}
	t.findByEmailCheckinReq = &dto.FindByEmailCheckInRequest{
		Email: t.checkin.Email,
	}
}

func (t *CheckInHandlerTest) TestCreateCheckinSuccess() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedResp := &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			Email:  t.checkin.Email,
			UserID: t.checkin.UserID,
			Event:  t.checkin.Event,
		},
	}

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(nil)
	checkinSvc.EXPECT().Create(t.createCheckinReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinBindError() {
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(nil, nil, t.logger)

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinValidationError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedError := []string{"error1", "error2"}

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(expectedError)
	context.EXPECT().BadRequestError("error1, error2")

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(nil)
	checkinSvc.EXPECT().Create(t.createCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.Create(context)

}

func (t *CheckInHandlerTest) TestFindByEmailCheckinSuccess() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedResp := &dto.FindByEmailCheckInResponse{
		CheckIns: t.checkins,
	}

	context.EXPECT().Bind(&dto.FindByEmailCheckInRequest{}).SetArg(0, *t.findByEmailCheckinReq)
	validator.EXPECT().Validate(t.findByEmailCheckinReq).Return(nil)
	checkinSvc.EXPECT().FindByEmail(t.findByEmailCheckinReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByEmailCheckinBindError() {
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(nil, nil, t.logger)

	context.EXPECT().Bind(&dto.FindByEmailCheckInRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByEmailCheckinValidationError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedError := []string{"error1", "error2"}

	context.EXPECT().Bind(&dto.FindByEmailCheckInRequest{}).SetArg(0, *t.findByEmailCheckinReq)
	validator.EXPECT().Validate(t.findByEmailCheckinReq).Return(expectedError)
	context.EXPECT().BadRequestError("error1, error2")

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByEmailCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	context.EXPECT().Bind(&dto.FindByEmailCheckInRequest{}).SetArg(0, *t.findByEmailCheckinReq)
	validator.EXPECT().Validate(t.findByEmailCheckinReq).Return(nil)
	checkinSvc.EXPECT().FindByEmail(t.findByEmailCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinSuccess() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedResp := &dto.FindByUserIdCheckInResponse{
		CheckIns: t.checkins,
	}

	context.EXPECT().Bind(&dto.FindByUserIdCheckInRequest{}).SetArg(0, *t.findByUserIdCheckinReq)
	validator.EXPECT().Validate(t.findByUserIdCheckinReq).Return(nil)
	checkinSvc.EXPECT().FindByUserID(t.findByUserIdCheckinReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByUserID(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinBindError() {
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(nil, nil, t.logger)

	context.EXPECT().Bind(&dto.FindByUserIdCheckInRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.FindByUserID(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinValidationError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	expectedError := []string{"error1", "error2"}

	context.EXPECT().Bind(&dto.FindByUserIdCheckInRequest{}).SetArg(0, *t.findByUserIdCheckinReq)
	validator.EXPECT().Validate(t.findByUserIdCheckinReq).Return(expectedError)
	context.EXPECT().BadRequestError("error1, error2")

	handler.FindByUserID(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := checkin.NewHandler(checkinSvc, validator, t.logger)

	context.EXPECT().Bind(&dto.FindByUserIdCheckInRequest{}).SetArg(0, *t.findByUserIdCheckinReq)
	validator.EXPECT().Validate(t.findByUserIdCheckinReq).Return(nil)
	checkinSvc.EXPECT().FindByUserID(t.findByUserIdCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByUserID(context)
}

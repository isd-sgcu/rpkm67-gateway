package test

import (
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	checkinMock "github.com/isd-sgcu/rpkm67-gateway/mocks/checkin"
	contextMock "github.com/isd-sgcu/rpkm67-gateway/mocks/context"
	userMock "github.com/isd-sgcu/rpkm67-gateway/mocks/user"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
)

type CheckInHandlerTest struct {
	suite.Suite
	controller             *gomock.Controller
	logger                 *zap.Logger
	checkins               []*dto.CheckIn
	checkin                *dto.CheckIn
	user                   *dto.User
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
	t.user = &dto.User{
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
	}

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
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, userSvc, validator, t.logger)

	expectedResp := &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			Email:  t.checkin.Email,
			UserID: t.checkin.UserID,
			Event:  t.checkin.Event,
		},
		Firstname: t.user.Firstname,
		Lastname:  t.user.Lastname,
	}

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(nil)
	checkinSvc.EXPECT().Create(t.createCheckinReq).Return(expectedResp, nil)
	userSvc.EXPECT().FindOne(&dto.FindOneUserRequest{Id: t.checkin.UserID}).Return(&dto.FindOneUserResponse{User: t.user}, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinBindError() {
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(nil, nil, nil, t.logger)

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinValidationError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, validator, t.logger)

	expectedError := []string{"error1", "error2"}

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(expectedError)
	context.EXPECT().BadRequestError("error1, error2")

	handler.Create(context)
}

func (t *CheckInHandlerTest) TestCreateCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, validator, t.logger)

	context.EXPECT().Bind(&dto.CreateCheckInRequest{}).SetArg(0, *t.createCheckinReq)
	validator.EXPECT().Validate(t.createCheckinReq).Return(nil)
	checkinSvc.EXPECT().Create(t.createCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.Create(context)

}

func (t *CheckInHandlerTest) TestFindByEmailCheckinSuccess() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, nil, t.logger)

	expectedResp := &dto.FindByEmailCheckInResponse{
		CheckIns: t.checkins,
	}

	context.EXPECT().Param("email").Return(t.checkin.Email)
	checkinSvc.EXPECT().FindByEmail(t.findByEmailCheckinReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByEmailCheckInParamEmpty() {
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(nil, nil, nil, t.logger)

	context.EXPECT().Param("email").Return("")
	context.EXPECT().BadRequestError(apperror.BadRequestError("email should not be empty").Error())

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByEmailCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, nil, t.logger)

	context.EXPECT().Param("email").Return(t.checkin.Email)
	checkinSvc.EXPECT().FindByEmail(t.findByEmailCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByEmail(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinSuccess() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, nil, t.logger)

	expectedResp := &dto.FindByUserIdCheckInResponse{
		CheckIns: t.checkins,
	}

	context.EXPECT().Param("userId").Return(t.checkin.UserID)
	checkinSvc.EXPECT().FindByUserID(t.findByUserIdCheckinReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByUserID(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinParamEmpty() {
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(nil, nil, nil, t.logger)

	expectedErr := apperror.BadRequestError("user_id should not be empty").Error()

	context.EXPECT().Param("userId").Return("")
	context.EXPECT().BadRequestError(expectedErr)

	handler.FindByUserID(context)
}

func (t *CheckInHandlerTest) TestFindByUserIdCheckinServiceError() {
	checkinSvc := checkinMock.NewMockService(t.controller)
	context := contextMock.NewMockCtx(t.controller)
	handler := checkin.NewHandler(checkinSvc, nil, nil, t.logger)

	context.EXPECT().Param("userId").Return(t.checkin.UserID)
	checkinSvc.EXPECT().FindByUserID(t.findByUserIdCheckinReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByUserID(context)
}

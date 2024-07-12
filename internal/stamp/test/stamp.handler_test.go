package test

import (
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/stamp"
	ctxMock "github.com/isd-sgcu/rpkm67-gateway/mocks/context"
	stampMock "github.com/isd-sgcu/rpkm67-gateway/mocks/stamp"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type StampHandlerTest struct {
	suite.Suite
	controller           *gomock.Controller
	logger               *zap.Logger
	stamp                *dto.Stamp
	findByUserIdReq      *dto.FindByUserIdStampRequest
	stampByUserIdReq     *dto.StampByUserIdRequest
	stampByUserIdBodyReq *dto.StampByUserIdBodyRequest
}

func TestStampHandler(t *testing.T) {
	suite.Run(t, new(StampHandlerTest))
}

func (t *StampHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	t.stamp = &dto.Stamp{
		ID:     faker.UUIDDigit(),
		UserID: faker.UUIDDigit(),
		PointA: 1,
		PointB: 2,
		PointC: 3,
		PointD: 4,
		Stamp:  faker.Word(),
	}
	t.findByUserIdReq = &dto.FindByUserIdStampRequest{
		UserID: faker.UUIDDigit(),
	}
	t.stampByUserIdBodyReq = &dto.StampByUserIdBodyRequest{
		ActivityId: faker.Word(),
		PinCode:    faker.Word(),
	}
	t.stampByUserIdReq = &dto.StampByUserIdRequest{
		UserID:     faker.UUIDDigit(),
		ActivityId: t.stampByUserIdBodyReq.ActivityId,
		PinCode:    t.stampByUserIdBodyReq.PinCode,
	}
}

func (t *StampHandlerTest) TestFindByUserIdSuccess() {
	stampSvc := stampMock.NewMockService(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(stampSvc, nil, t.logger)

	expectedResp := &dto.FindByUserIdStampResponse{
		Stamp: t.stamp,
	}

	context.EXPECT().Param("userId").Return(t.findByUserIdReq.UserID)
	stampSvc.EXPECT().FindByUserId(t.findByUserIdReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByUserId(context)
}

func (t *StampHandlerTest) TestFindByUserIdMissingParamUserId() {
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(nil, nil, t.logger)

	context.EXPECT().Param("userId").Return("")
	context.EXPECT().BadRequestError("userId should not be empty")

	handler.FindByUserId(context)
}

func (t *StampHandlerTest) TestFindByUserIdServiceError() {
	stampSvc := stampMock.NewMockService(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(stampSvc, nil, t.logger)

	context.EXPECT().Param("userId").Return(t.findByUserIdReq.UserID)
	stampSvc.EXPECT().FindByUserId(t.findByUserIdReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByUserId(context)
}

func (t *StampHandlerTest) TestStampByUserIdSuccess() {
	stampSvc := stampMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(stampSvc, validator, t.logger)

	expectedResp := &dto.StampByUserIdResponse{
		Stamp: t.stamp,
	}

	context.EXPECT().Param("userId").Return(t.stampByUserIdReq.UserID)
	context.EXPECT().Bind(&dto.StampByUserIdBodyRequest{}).SetArg(0, *t.stampByUserIdBodyReq)
	validator.EXPECT().Validate(t.stampByUserIdBodyReq).Return(nil)

	stampSvc.EXPECT().StampByUserId(t.stampByUserIdReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.StampByUserId(context)
}

func (t *StampHandlerTest) TestStampByUserIdMissingParamUserId() {
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(nil, nil, t.logger)

	context.EXPECT().Param("userId").Return("")
	context.EXPECT().BadRequestError("userId should not be empty")

	handler.StampByUserId(context)
}

func (t *StampHandlerTest) TestStampByUserIdBindError() {
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(nil, nil, t.logger)

	context.EXPECT().Param("userId").Return(t.stampByUserIdReq.UserID)
	context.EXPECT().Bind(&dto.StampByUserIdBodyRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.StampByUserId(context)
}

func (t *StampHandlerTest) TestStampByUserIdServiceError() {
	stampSvc := stampMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := ctxMock.NewMockCtx(t.controller)
	handler := stamp.NewHandler(stampSvc, validator, t.logger)

	context.EXPECT().Param("userId").Return(t.stampByUserIdReq.UserID)
	context.EXPECT().Bind(&dto.StampByUserIdBodyRequest{}).SetArg(0, *t.stampByUserIdBodyReq)
	validator.EXPECT().Validate(t.stampByUserIdBodyReq).Return(nil)
	stampSvc.EXPECT().StampByUserId(t.stampByUserIdReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.StampByUserId(context)
}

package test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/baan"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	baanMock "github.com/isd-sgcu/rpkm67-gateway/mocks/baan"
	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BaanHandlerTest struct {
	suite.Suite
	controller     *gomock.Controller
	logger         *zap.Logger
	Baans          []*dto.Baan
	Baan           *dto.Baan
	FindAllBaanReq *dto.FindAllBaanRequest
	FindOneBaanReq *dto.FindOneBaanRequest
	Err            *apperrors.AppError
	ParamMock      string
}

func TestBaanHandler(t *testing.T) {
	suite.Run(t, new(BaanHandlerTest))
}

func (t *BaanHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	baansProto := MockBaansProto()
	baanProto := baansProto[0]

	t.Baans = baan.ProtoToDtoList(baansProto)
	t.Baan = baan.ProtoToDto(baanProto)

	t.FindAllBaanReq = &dto.FindAllBaanRequest{}
	t.FindOneBaanReq = &dto.FindOneBaanRequest{
		Id: t.Baan.Id,
	}

	t.ParamMock = t.Baan.Id
}

func (t *BaanHandlerTest) TestFindAllBaan() {
	expectedResp := &dto.FindAllBaanResponse{
		Baans: t.Baans,
	}

	controller := gomock.NewController(t.T())

	baanSvc := baanMock.NewMockService(controller)
	validator := validatorMock.NewMockDtoValidator(controller)
	context := routerMock.NewMockContext(controller)

	baanSvc.EXPECT().FindAllBaan(t.FindAllBaanReq).Return(expectedResp, t.Err)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler := baan.NewHandler(baanSvc, validator, t.logger)
	handler.FindAllBaan(context)
}

func (t *BaanHandlerTest) TestFindOneBaan() {
	expectedResp := &dto.FindOneBaanResponse{
		Baan: t.Baan,
	}

	controller := gomock.NewController(t.T())

	baanSvc := baanMock.NewMockService(controller)
	validator := validatorMock.NewMockDtoValidator(controller)
	context := routerMock.NewMockContext(controller)

	context.EXPECT().Param("id").Return(t.ParamMock)
	validator.EXPECT().Validate(t.FindOneBaanReq).Return(nil)
	baanSvc.EXPECT().FindOneBaan(t.FindOneBaanReq).Return(expectedResp, t.Err)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler := baan.NewHandler(baanSvc, validator, t.logger)
	handler.FindOneBaan(context)
}

func (t *BaanHandlerTest) TearDownTest() {
	t.controller.Finish()
}

package test

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/isd-sgcu/rpkm67-gateway/apperror"
// 	"github.com/isd-sgcu/rpkm67-gateway/internal/baan"
// 	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
// 	baanMock "github.com/isd-sgcu/rpkm67-gateway/mocks/baan"
// 	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
// 	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
// 	"github.com/stretchr/testify/suite"
// 	"go.uber.org/zap"
// )

// type BaanHandlerTest struct {
// 	suite.Suite
// 	controller     *gomock.Controller
// 	logger         *zap.Logger
// 	Baans          []*dto.Baan
// 	Baan           *dto.Baan
// 	FindAllBaanReq *dto.FindAllBaanRequest
// 	FindOneBaanReq *dto.FindOneBaanRequest
// 	Err            *apperror.AppError
// 	ParamMock      string
// }

// func TestBaanHandler(t *testing.T) {
// 	suite.Run(t, new(BaanHandlerTest))
// }

// func (t *BaanHandlerTest) SetupTest() {
// 	t.controller = gomock.NewController(t.T())
// 	t.logger = zap.NewNop()

// 	baansProto := MockBaansProto()
// 	baanProto := baansProto[0]

// 	t.Baans = baan.ProtoToDtoList(baansProto)
// 	t.Baan = baan.ProtoToDto(baanProto)

// 	t.FindAllBaanReq = &dto.FindAllBaanRequest{}
// 	t.FindOneBaanReq = &dto.FindOneBaanRequest{
// 		Id: t.Baan.Id,
// 	}

// 	t.ParamMock = t.Baan.Id
// }

// func (t *BaanHandlerTest) TestFindAllBaanSuccess() {
// 	baanSvc := baanMock.NewMockService(t.controller)
// 	validator := validatorMock.NewMockDtoValidator(t.controller)
// 	context := routerMock.NewMockContext(t.controller)
// 	handler := baan.NewHandler(baanSvc, validator, t.logger)

// 	expectedResp := &dto.FindAllBaanResponse{
// 		Baans: t.Baans,
// 	}

// 	baanSvc.EXPECT().FindAllBaan(t.FindAllBaanReq).Return(expectedResp, t.Err)
// 	context.EXPECT().JSON(http.StatusOK, expectedResp)

// 	handler.FindAllBaan(context)
// }

// func (t *BaanHandlerTest) TestFindOneBaanSuccess() {
// 	baanSvc := baanMock.NewMockService(t.controller)
// 	validator := validatorMock.NewMockDtoValidator(t.controller)
// 	context := routerMock.NewMockContext(t.controller)
// 	handler := baan.NewHandler(baanSvc, validator, t.logger)

// 	expectedResp := &dto.FindOneBaanResponse{
// 		Baan: t.Baan,
// 	}

// 	context.EXPECT().Param("id").Return(t.ParamMock)
// 	validator.EXPECT().Validate(t.FindOneBaanReq).Return(nil)
// 	baanSvc.EXPECT().FindOneBaan(t.FindOneBaanReq).Return(expectedResp, t.Err)
// 	context.EXPECT().JSON(http.StatusOK, expectedResp)

// 	handler.FindOneBaan(context)
// }

// func (t *BaanHandlerTest) TearDownTest() {
// 	t.controller.Finish()
// }

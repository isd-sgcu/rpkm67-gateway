package test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/selection"
	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
	selectionMock "github.com/isd-sgcu/rpkm67-gateway/mocks/selection"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SelectionHandlerTest struct {
	suite.Suite
	controller                *gomock.Controller
	logger                    *zap.Logger
	Selections                []*dto.Selection
	Selection                 *dto.Selection
	CreateSelectionReq        *dto.CreateSelectionRequest
	FindByGroupIdSelectionReq *dto.FindByGroupIdSelectionRequest
	UpdateSelectionReq        *dto.UpdateSelectionRequest
}

func TestSelectionHandler(t *testing.T) {
	suite.Run(t, new(SelectionHandlerTest))
}

func (t *SelectionHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()

	selectionsProto := MockSelectionsProto()
	selectionProto := selectionsProto[0]

	t.Selections = selection.ProtoToDtoList(selectionsProto)
	t.Selection = selection.ProtoToDto(selectionProto)

	t.CreateSelectionReq = &dto.CreateSelectionRequest{
		GroupId: t.Selection.GroupId,
		BaanIds: t.Selection.BaanIds,
	}
	t.FindByGroupIdSelectionReq = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.Selection.GroupId,
	}
	t.UpdateSelectionReq = &dto.UpdateSelectionRequest{
		Selection: selection.ProtoToDto(selectionsProto[1]),
	}
}

func (t *SelectionHandlerTest) TestCreateSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedResp := &dto.CreateSelectionResponse{
		Selection: t.Selection,
	}

	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).SetArg(0, *t.CreateSelectionReq)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionBindError() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	context.EXPECT().Bind(&dto.CreateSelectionRequest{}).SetArg(0, *t.CreateSelectionReq)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedResp := &dto.FindByGroupIdSelectionResponse{
		Selection: t.Selection,
	}

	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupIdSelection(t.FindByGroupIdSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionUrlParamEmpty() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	expectedErr := apperror.BadRequestError("url parameter 'id' not found")

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupIdSelection(t.FindByGroupIdSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedResp := &dto.UpdateSelectionResponse{
		Success: true,
	}

	context.EXPECT().Bind(&dto.UpdateSelectionRequest{}).SetArg(0, *t.UpdateSelectionReq)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)
	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionBindError() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	context.EXPECT().Bind(&dto.UpdateSelectionRequest{}).Return(apperror.BadRequest)
	context.EXPECT().BadRequestError(apperror.BadRequest.Error())

	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionServiceError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	context.EXPECT().Bind(&dto.UpdateSelectionRequest{}).SetArg(0, *t.UpdateSelectionReq)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)
	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.UpdateSelection(context)
}

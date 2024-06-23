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
	Err                       *apperror.AppError
	QueriesMock               map[string]string
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

	t.CreateSelectionReq = &dto.CreateSelectionRequest{}
	t.FindByGroupIdSelectionReq = &dto.FindByGroupIdSelectionRequest{
		GroupId: t.Selection.GroupId,
	}
	t.UpdateSelectionReq = &dto.UpdateSelectionRequest{}
}

func (t *SelectionHandlerTest) TestCreateSelectionSuccess() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedResp := &dto.CreateSelectionResponse{
		Selection: t.Selection,
	}

	context.EXPECT().Bind(t.CreateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusCreated, expectedResp)

	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionInvalidArgument() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	expectedErr := apperror.BadRequestError("invalid argument")

	context.EXPECT().Bind(t.CreateSelectionReq).Return(expectedErr)
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler.CreateSelection(context)
}

func (t *SelectionHandlerTest) TestCreateSelectionInternalError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedErr := apperror.InternalServerError("internal error")

	context.EXPECT().Bind(t.CreateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.CreateSelectionReq).Return(nil)
	selectionSvc.EXPECT().CreateSelection(t.CreateSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

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

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionInvalidArgument() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	expectedErr := apperror.BadRequestError("url parameter 'id' not found")

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler.FindByGroupIdSelection(context)
}

func (t *SelectionHandlerTest) TestFindByStudentIdSelectionInternalError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedErr := apperror.InternalServerError("internal error")

	context.EXPECT().Param("id").Return(t.Selection.GroupId)
	validator.EXPECT().Validate(t.FindByGroupIdSelectionReq).Return(nil)
	selectionSvc.EXPECT().FindByGroupIdSelection(t.FindByGroupIdSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

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

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)
	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionInvalidArgument() {
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(nil, nil, t.logger)

	expectedErr := apperror.BadRequestError("invalid argument")

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(expectedErr)
	context.EXPECT().BadRequestError(expectedErr.Error())

	handler.UpdateSelection(context)
}

func (t *SelectionHandlerTest) TestUpdateSelectionInternalError() {
	selectionSvc := selectionMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := selection.NewHandler(selectionSvc, validator, t.logger)

	expectedErr := apperror.InternalServerError("internal error")

	context.EXPECT().Bind(t.UpdateSelectionReq).Return(nil)
	validator.EXPECT().Validate(t.UpdateSelectionReq).Return(nil)
	selectionSvc.EXPECT().UpdateSelection(t.UpdateSelectionReq).Return(nil, expectedErr)
	context.EXPECT().ResponseError(expectedErr)

	handler.UpdateSelection(context)
}

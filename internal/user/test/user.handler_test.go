package test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/user"
	rpkm_constant "github.com/isd-sgcu/rpkm67-gateway/constant"
	routerMock "github.com/isd-sgcu/rpkm67-gateway/mocks/router"
	userMock "github.com/isd-sgcu/rpkm67-gateway/mocks/user"
	validatorMock "github.com/isd-sgcu/rpkm67-gateway/mocks/validator"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UserHandlerTest struct {
	suite.Suite
	controller               *gomock.Controller
	logger                   *zap.Logger
	User                     *dto.User
	FindOneUserRequest       *dto.FindOneUserRequest
	UpdateUserProfileBody    *dto.UpdateUserProfileBody
	UpdateUserProfileRequest *dto.UpdateUserProfileRequest
	UpdateUserPictureRequest *dto.UpdateUserPictureRequest
	DecomposedFile           *dto.DecomposedFile
	contentType              map[string]struct{}
	maxSizeAccept            int64
	Err                      *apperror.AppError
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

func (t *UserHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
	t.contentType = rpkm_constant.AllowedContentType
	t.maxSizeAccept = 20

	userProto := MockUserProto()

	t.User = user.ProtoToDto(userProto)

	t.DecomposedFile = &dto.DecomposedFile{
		Filename: "test.jpg",
		Data:     []byte{},
	}

	t.FindOneUserRequest = &dto.FindOneUserRequest{
		Id: t.User.Id,
	}

	t.UpdateUserProfileBody = &dto.UpdateUserProfileBody{
		Nickname:    t.User.Nickname,
		Title:       t.User.Title,
		Firstname:   t.User.Firstname,
		Lastname:    t.User.Lastname,
		Year:        t.User.Year,
		Faculty:     t.User.Faculty,
		Tel:         t.User.Tel,
		ParentTel:   t.User.ParentTel,
		Parent:      t.User.Parent,
		FoodAllergy: t.User.FoodAllergy,
		DrugAllergy: t.User.DrugAllergy,
		Illness:     t.User.Illness,
		PhotoKey:    t.User.PhotoKey,
		PhotoUrl:    t.User.PhotoUrl,
		Baan:        t.User.Baan,
		ReceiveGift: t.User.ReceiveGift,
		GroupId:     t.User.GroupId,
	}

	t.UpdateUserProfileRequest = &dto.UpdateUserProfileRequest{
		Id:	t.User.Id,
		Nickname:    t.User.Nickname,
		Title:       t.User.Title,
		Firstname:   t.User.Firstname,
		Lastname:    t.User.Lastname,
		Year:        t.User.Year,
		Faculty:     t.User.Faculty,
		Tel:         t.User.Tel,
		ParentTel:   t.User.ParentTel,
		Parent:      t.User.Parent,
		FoodAllergy: t.User.FoodAllergy,
		DrugAllergy: t.User.DrugAllergy,
		Illness:     t.User.Illness,
		PhotoKey:    t.User.PhotoKey,
		PhotoUrl:    t.User.PhotoUrl,
		Baan:        t.User.Baan,
		ReceiveGift: t.User.ReceiveGift,
		GroupId:     t.User.GroupId,
	}

	t.UpdateUserPictureRequest = &dto.UpdateUserPictureRequest{
		Id:   t.User.Id,
		File: t.DecomposedFile,
	}

}

func (t *UserHandlerTest) TestFindOneUserSuccess() {
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, 0, nil, validator, t.logger)

	expectedResp := &dto.FindOneUserResponse{
		User: t.User,
	}

	context.EXPECT().Param("id").Return(t.User.Id)
	userSvc.EXPECT().FindOne(t.FindOneUserRequest).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.FindOne(context)
}

func (t *UserHandlerTest) TestFindOneUserParamEmpty() {
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(nil, 0, nil, validator, t.logger)

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError("url parameter 'id' not found")

	handler.FindOne(context)
}

func (t *UserHandlerTest) TestFindOneUserServiceError() {
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, 0, nil, validator, t.logger)

	context.EXPECT().Param("id").Return(t.User.Id)
	userSvc.EXPECT().FindOne(t.FindOneUserRequest).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.FindOne(context)
}

func (t *UserHandlerTest) TestUpdateProfileSuccess() {
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, 0, nil, validator, t.logger)

	expectedResp := &dto.UpdateUserProfileResponse{
		User: t.User,
	}

	context.EXPECT().Param("id").Return(t.User.Id)
	context.EXPECT().Bind(&dto.UpdateUserProfileBody{}).SetArg(0, *t.UpdateUserProfileBody)
	validator.EXPECT().Validate(t.UpdateUserProfileBody).Return(nil)
	userSvc.EXPECT().UpdateProfile(t.UpdateUserProfileRequest).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.UpdateProfile(context)

}

func (t *UserHandlerTest) TestUpdateProfileParamEmpty() {
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(nil, 0, nil, nil, t.logger)

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError("url parameter 'id' not found")

	handler.UpdateProfile(context)
}

func (t *UserHandlerTest) TestUpdateProfileServiceError() {
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, int(t.maxSizeAccept), t.contentType, validator, t.logger)

	context.EXPECT().Param("id").Return(t.User.Id)
	context.EXPECT().Bind(&dto.UpdateUserProfileBody{}).SetArg(0, *t.UpdateUserProfileBody)
	validator.EXPECT().Validate(t.UpdateUserProfileBody).Return(nil)
	userSvc.EXPECT().UpdateProfile(t.UpdateUserProfileRequest).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.UpdateProfile(context)
}

func (t *UserHandlerTest) TestUpdatePictureSuccess() {
	userSvc := userMock.NewMockService(t.controller)
	validator := validatorMock.NewMockDtoValidator(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, int(t.maxSizeAccept), t.contentType, validator, t.logger)

	expectedResp := &dto.UpdateUserPictureResponse{
		User: t.User,
	}

	context.EXPECT().Param("id").Return(t.User.Id)
	context.EXPECT().FormFile("picture", t.contentType, t.maxSizeAccept).Return(t.UpdateUserPictureRequest.File, nil)
	userSvc.EXPECT().UpdatePicture(t.UpdateUserPictureRequest).Return(expectedResp, nil)
	context.EXPECT().JSON(http.StatusOK, expectedResp)

	handler.UpdatePicture(context)
}

func (t *UserHandlerTest) TestUpdatePictureParamEmpty() {
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(nil, int(t.maxSizeAccept), nil, nil, t.logger)

	context.EXPECT().Param("id").Return("")
	context.EXPECT().BadRequestError("url parameter 'id' not found")

	handler.UpdatePicture(context)
}

func (t *UserHandlerTest) TestUpdatePictureServiceError() {
	userSvc := userMock.NewMockService(t.controller)
	context := routerMock.NewMockContext(t.controller)
	handler := user.NewHandler(userSvc, int(t.maxSizeAccept), t.contentType, nil, t.logger)

	context.EXPECT().Param("id").Return(t.User.Id)
	context.EXPECT().FormFile("picture", t.contentType, t.maxSizeAccept).Return(t.UpdateUserPictureRequest.File, nil)
	userSvc.EXPECT().UpdatePicture(t.UpdateUserPictureRequest).Return(nil, apperror.InternalServer)
	context.EXPECT().ResponseError(apperror.InternalServer)

	handler.UpdatePicture(context)
}

func (t *UserHandlerTest) TearDownTest() {
	t.controller.Finish()
}

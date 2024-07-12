package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UserHandlerTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

func (t *UserHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *UserHandlerTest) TestFindOneUserSuccess() {

}

func (t *UserHandlerTest) TearDownTest() {
	t.controller.Finish()
}

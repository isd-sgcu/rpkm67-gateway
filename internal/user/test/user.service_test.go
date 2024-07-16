package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UserServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *UserServiceTest) TestFindOneUserSuccess() {

}

func (t *UserServiceTest) TearDownTest() {
	t.controller.Finish()
}

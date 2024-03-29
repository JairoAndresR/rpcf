package managers

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"rpcf/accounts/ports"
	"rpcf/accounts/ports/mocks"
	"rpcf/app/accounts/adapters/repositories/fixtures"
	"testing"
)

type registerManagerSuite struct {
	suite.Suite
	accManager *mocks.AccountManager
	token      *mocks.TokenManager
	manager    ports.RegisterManager
}

func TestRegisterManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(registerManagerSuite))
}

func (s *registerManagerSuite) SetupSuite() {
	s.accManager = &mocks.AccountManager{}
	s.token = &mocks.TokenManager{}

	s.manager = newRegisterManager(s.accManager, s.token)
}

func (s *registerManagerSuite) TestRegisterError() {
	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Create", mock.Anything).
		Return(nil, errors.New("account already exist")).
		Once()

	_, token, err := s.manager.Register(acc)
	s.NotNil(err)
	s.Empty(token)
}

func (s *registerManagerSuite) TestRegisterTokenError() {
	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Create", mock.Anything).
		Return(acc, nil).
		Once()

	s.token.On("Generate", mock.Anything).
		Return("", errors.New(TokenError)).
		Once()

	_, token, err := s.manager.Register(acc)
	s.NotNil(err)
	s.Equal(TokenError, err.Error())
	s.Empty(token)
}

func (s *registerManagerSuite) TestRegister() {
	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Create", mock.Anything).
		Return(acc, nil).
		Once()

	tokenExpected := "1234545"
	s.token.On("Generate", mock.Anything).
		Return(tokenExpected, nil).
		Once()

	_, token, err := s.manager.Register(acc)
	s.Nil(err)
	s.Equal(tokenExpected, token)
}

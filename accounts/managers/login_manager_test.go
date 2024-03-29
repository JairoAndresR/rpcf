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

type loginManagerSuite struct {
	suite.Suite
	accManager *mocks.AccountManager
	token      *mocks.TokenManager

	manager ports.LoginManager
}

func TestLoginManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(loginManagerSuite))
}

func (s *loginManagerSuite) SetupSuite() {
	s.accManager = &mocks.AccountManager{}
	s.token = &mocks.TokenManager{}
	s.manager = newLoginManager(s.accManager, s.token)
}

func (s *loginManagerSuite) TestLoginAccountNotFound() {
	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Login", mock.Anything).
		Return(nil, errors.New(UserNotFound)).
		Once()

	_, token, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(UserNotFound, err.Error())
	s.Empty(token)
}

func (s *loginManagerSuite) TestLoginAccountTokenError() {

	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Login", mock.Anything).
		Return(acc, nil).
		Once()

	s.token.On("Generate", mock.Anything).
		Return("", errors.New(TokenError)).
		Once()

	_, token, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(TokenError, err.Error())
	s.Empty(token)
}

func (s *loginManagerSuite) TestLogin() {

	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Login", mock.Anything).
		Return(acc, nil).
		Once()

	tokenExpected := "12345"
	s.token.On("Generate", mock.Anything).
		Return(tokenExpected, nil).
		Once()

	_, token, err := s.manager.Login(acc)
	s.Nil(err)
	s.Equal(tokenExpected, token)
}

package managers

import (
	"github.com/stretchr/testify/suite"
	"rpcf/accounts/ports"
	"rpcf/accounts/ports/fixtures"
	"testing"
)

type tokeManagerSuite struct {
	suite.Suite
	manager ports.TokenManager
}

func TestTokenManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(tokeManagerSuite))
}

func (s *tokeManagerSuite) SetupSuite() {
	s.manager = newTokenManager()
}

func (s *tokeManagerSuite) TestGenerate() {
	acc := fixtures.GetAccount()
	token, err := s.manager.Generate(acc)
	s.Nil(err)
	s.NotEmpty(token)
}

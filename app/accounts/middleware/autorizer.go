package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/accounts/ports"
	"rpcf/app/accounts/adapters/handlers"
	"rpcf/core"
	"strings"
)

func init() {
	err := core.Injector.Provide(NewAuthorizer)
	core.CheckInjection(err, "NewAuthorizer")
}

const (
	ApiTokenRequired = "api token required"
	InvalidApiToken  = "invalid API token"
	AdminRole        = "admin"
)

type Authorizer struct {
	manager ports.TokenManager
}

func NewAuthorizer(manager ports.TokenManager) *Authorizer {

	return &Authorizer{
		manager: manager,
	}
}

func (a *Authorizer) Token() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := a.AuthorizeUser(c)
		if err != nil {
			return
		}
		c.Next()
	}
}

func (a *Authorizer) AuthorizeAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := a.AuthorizeUser(c)
		if err != nil {
			return
		}
		value, exist := c.Get("role")
		if !exist {
			err := errors.New(ApiTokenRequired)
			handlers.AbortWithError(c, http.StatusUnauthorized, err)
			return
		}

		if value != AdminRole {
			err := errors.New(ApiTokenRequired)
			handlers.AbortWithError(c, http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}

func (a *Authorizer) AuthorizeUser(c *gin.Context) error {
	token, ok := a.getToken(c.Request.Header)

	if !ok {
		err := errors.New(ApiTokenRequired)
		handlers.AbortWithError(c, http.StatusUnauthorized, err)
		return err
	}
	isValid, err := a.manager.Validate(token)

	if err != nil || !isValid {
		err = errors.New(InvalidApiToken)
		handlers.AbortWithError(c, http.StatusUnauthorized, err)
		return err
	}

	account, err := a.manager.GetMetadata(token)
	if err != nil || !isValid {
		err = errors.New(InvalidApiToken)
		handlers.AbortWithError(c, http.StatusUnauthorized, err)
		return err
	}

	c.Set("userId", account.ID)
	c.Set("names", account.Names)
	c.Set("email", account.Email)
	return nil
}

func (a *Authorizer) getToken(headers http.Header) (string, bool) {
	token := headers.Get("Authorization")
	values := strings.Split(token, "Bearer ")
	if token == "" || len(values) != 2 {
		return "", false
	}
	return values[1], true
}

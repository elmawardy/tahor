package authapi

import (
	"context"
	"log"

	auth "github.com/elmawardy/tahor/auth/gen/auth"
	"github.com/elmawardy/tahor/auth/models"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	logger *log.Logger
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger) auth.Service {
	return &authsrvc{logger}
}

// Creates a valid JWT
func (s *authsrvc) Signin(ctx context.Context, p *auth.SigninPayload) (res *auth.Creds, err error) {

	u := &models.User{}
	err = u.ValidatePassword(p.Email, p.Password)
	if err != nil {
		return
	}

	user := &models.User{}
	token, err := user.GenerateJWTFromEmail(p.Email)
	if err != nil {
		return
	}

	res = &auth.Creds{}
	res.JWT = token
	return
}

func (s *authsrvc) Register(ctx context.Context, p *auth.RegisterPayload) (res string, err error) {
	u := &models.User{}
	u.Email = p.Email
	u.Password = p.Password
	if err != nil {
		return
	}
	u.Name = *p.Name

	err = u.Register()

	if err != nil {
		return
	}

	res = "success"
	return
}

func (s *authsrvc) Signout(ctx context.Context, p *auth.SignoutPayload) (res string, err error) {
	u := models.User{}
	err = u.Logout(p.JWT)
	res = "success"
	return
}

func (s *authsrvc) Getbasicinfo(ctx context.Context, p *auth.GetbasicinfoPayload) (res *auth.UserBasicInfo, err error) {
	u := models.User{}
	info, err := u.GetBasicInfo(p.JWT)
	if err != nil {
		return
	}

	res = &auth.UserBasicInfo{}
	res.Name = &info.Name
	return
}

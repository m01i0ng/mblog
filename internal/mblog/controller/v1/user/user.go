package user

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/pkg/auth"
)

type UserController struct {
	a *auth.Authz
	b biz.IBiz
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{
		a: a,
		b: biz.New(ds),
	}
}

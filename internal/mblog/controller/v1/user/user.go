package user

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz"
	"github.com/m01i0ng/mblog/internal/mblog/store"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.New(ds)}
}

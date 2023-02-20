package biz

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz/user"
	"github.com/m01i0ng/mblog/internal/mblog/store"
)

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IStore
}

func New(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

var _ IBiz = (*biz)(nil)

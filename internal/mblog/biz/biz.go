// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package biz

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz/post"
	"github.com/m01i0ng/mblog/internal/mblog/biz/user"
	"github.com/m01i0ng/mblog/internal/mblog/store"
)

type IBiz interface {
	Users() user.UserBiz
	Posts() post.PostBiz
}

type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func New(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

func (b *biz) Posts() post.PostBiz {
	return post.New(b.ds)
}

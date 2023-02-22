// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package user

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/pkg/auth"
	pb "github.com/m01i0ng/mblog/pkg/proto/mblog/v1"
)

type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedMBlogServer
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{
		a: a,
		b: biz.New(ds),
	}
}

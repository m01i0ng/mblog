// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/known"
	"github.com/m01i0ng/mblog/internal/pkg/log"
)

type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

func Authz(a Auther) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString(known.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method

		log.Debugw("Build authorize context", "sub", sub, "obj", obj, "act", act)
		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			core.WriteResponse(c, errno.ErrUnauthorized, nil)
			c.Abort()
			return
		}
	}
}

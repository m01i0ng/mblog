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
	"github.com/m01i0ng/mblog/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}

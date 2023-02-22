// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package mblog

import (
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/mblog/controller/v1/user"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/internal/pkg/middleware"
	"github.com/m01i0ng/mblog/pkg/auth"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz func called")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)

	v1 := g.Group("/v1")

	{
		userV1 := v1.Group("/users")
		{
			userV1.POST("", uc.Create)
			userV1.PUT(":name/change-password", uc.ChangePassword)
			userV1.Use(middleware.Authn(), middleware.Authz(authz))
			userV1.GET(":name", uc.Get)
		}
	}

	return nil
}

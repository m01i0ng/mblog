// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/log"
)

func (uc *UserController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete user func called")
	username := c.Param("name")

	if err := uc.b.Users().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := uc.a.RemoveNamedPolicy("p", username, "", ""); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}

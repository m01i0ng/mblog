// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
)

func (uc *UserController) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password func called")

	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := uc.b.Users().ChangePassword(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}

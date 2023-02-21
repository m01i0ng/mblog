package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

func (uc *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create func called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := uc.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := uc.a.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, defaultMethods); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}

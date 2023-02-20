package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/mblog/biz"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.New(ds)}
}

func (u *UserController) Create(c *gin.Context) {
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

	if err := u.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}

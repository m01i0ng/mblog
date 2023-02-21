package user

import (
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/log"
)

func (uc *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user func called")

	user, err := uc.b.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, user)
}

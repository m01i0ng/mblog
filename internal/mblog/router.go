package mblog

import (
  "github.com/gin-gonic/gin"
  "github.com/m01i0ng/mblog/internal/mblog/controller/v1/user"
  "github.com/m01i0ng/mblog/internal/mblog/store"
  "github.com/m01i0ng/mblog/internal/pkg/core"
  "github.com/m01i0ng/mblog/internal/pkg/errno"
  "github.com/m01i0ng/mblog/internal/pkg/log"
)

func installRouters(g *gin.Engine) error {
  g.NoRoute(func(c *gin.Context) {
    core.WriteResponse(c, errno.ErrPageNotFound, nil)
  })

  g.GET("/healthz", func(c *gin.Context) {
    log.C(c).Infow("Healthz func called")
    core.WriteResponse(c, nil, map[string]string{"status": "ok"})
  })

  uc := user.New(store.S)

  v1 := g.Group("/v1")

  {
    userV1 := v1.Group("/users")
    {
      userV1.POST("", uc.Create)
    }
  }

  return nil
}

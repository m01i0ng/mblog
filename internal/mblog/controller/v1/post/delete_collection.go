package post

import (
	"github.com/gin-gonic/gin"
	"github.com/m01i0ng/mblog/internal/pkg/core"
	"github.com/m01i0ng/mblog/internal/pkg/known"
	"github.com/m01i0ng/mblog/internal/pkg/log"
)

func (p *PostController) DeleteCollection(c *gin.Context) {
	log.C(c).Infow("Batch delete post func called")

	postIDs := c.QueryArray("postID")
	if err := p.b.Posts().DeleteCollection(c, c.GetString(known.XUsernameKey), postIDs); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}

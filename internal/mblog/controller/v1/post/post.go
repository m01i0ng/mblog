package post

import (
	"github.com/m01i0ng/mblog/internal/mblog/biz"
	"github.com/m01i0ng/mblog/internal/mblog/store"
)

type PostController struct {
	b biz.IBiz
}

func New(ds store.IStore) *PostController {
	return &PostController{b: biz.New(ds)}
}

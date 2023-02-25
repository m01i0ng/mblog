package post

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/internal/pkg/model"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
	"gorm.io/gorm"
)

type PostBiz interface {
	Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
	Update(ctx context.Context, username, postID string, request *v1.UpdatePostRequest) error
	Delete(ctx context.Context, username, postID string) error
	DeleteCollection(ctx context.Context, username string, postIDs []string) error
	Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error)
	List(ctx context.Context, username string, offset, limit int) (*v1.ListPostResponse, error)
}

type postBiz struct {
	ds store.IStore
}

func New(ds store.IStore) *postBiz {
	return &postBiz{ds: ds}
}

var _ PostBiz = (*postBiz)(nil)

func (p *postBiz) Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	var postM model.PostM
	_ = copier.Copy(&postM, r)
	postM.Username = username

	if err := p.ds.Posts().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &v1.CreatePostResponse{PostID: postM.PostID}, nil
}

func (p *postBiz) Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error {
	postM, err := p.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		return err
	}

	if r.Title != nil {
		postM.Title = *r.Title
	}
	if r.Content != nil {
		postM.Content = *r.Content
	}

	if err = p.ds.Posts().Update(ctx, postM); err != nil {
		return err
	}
	return nil
}

func (p *postBiz) Delete(ctx context.Context, username, postID string) error {
	if err := p.ds.Posts().Delete(ctx, username, []string{postID}); err != nil {
		return err
	}

	return nil
}

func (p *postBiz) DeleteCollection(ctx context.Context, username string, postIDs []string) error {
	if err := p.ds.Posts().Delete(ctx, username, postIDs); err != nil {
		return err
	}
	return nil
}

func (p *postBiz) Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error) {
	post, err := p.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}
		return nil, err
	}

	var resp v1.GetPostResponse
	_ = copier.Copy(&resp, post)

	resp.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = post.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, err
}

func (p *postBiz) List(ctx context.Context, username string, offset, limit int) (*v1.ListPostResponse, error) {
	count, list, err := p.ds.Posts().List(ctx, username, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("Failed to load posts from storage", "err", err)
		return nil, err
	}

	posts := make([]*v1.PostInfo, 0, len(list))
	for _, item := range list {
		post := item
		posts = append(posts, &v1.PostInfo{
			Username:  post.Username,
			PostID:    post.PostID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.ListPostResponse{TotalCount: count, Posts: posts}, nil
}

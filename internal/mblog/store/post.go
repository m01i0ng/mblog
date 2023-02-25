package store

import (
	"context"
	"errors"

	"github.com/m01i0ng/mblog/internal/pkg/model"
	"gorm.io/gorm"
)

type PostStore interface {
	Create(ctx context.Context, post *model.PostM) error
	Get(ctx context.Context, username, postID string) (*model.PostM, error)
	Update(ctx context.Context, post *model.PostM) error
	List(ctx context.Context, username string, offset, limit int) (int64, []*model.PostM, error)
	Delete(ctx context.Context, username string, postIDs []string) error
}

type posts struct {
	db *gorm.DB
}

func newPosts(db *gorm.DB) *posts {
	return &posts{db: db}
}

var _ PostStore = (*posts)(nil)

func (p *posts) Create(ctx context.Context, post *model.PostM) error {
	return p.db.Create(&post).Error
}

func (p *posts) Get(ctx context.Context, username, postID string) (*model.PostM, error) {
	var post model.PostM
	if err := p.db.Where("username = ? and postID = ?", username, postID).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *posts) Update(ctx context.Context, post *model.PostM) error {
	return p.db.Save(post).Error
}

func (p *posts) List(ctx context.Context, username string, offset, limit int) (count int64, posts []*model.PostM, err error) {
	err = p.db.Where("username = ?", username).Offset(offset).Limit(defaultLimit(limit)).Order("id desc").Find(&posts).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
	return
}

func (p *posts) Delete(ctx context.Context, username string, postIDs []string) error {
	err := p.db.Where("username = ? and postID in (?)", username, postIDs).Delete(&model.PostM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

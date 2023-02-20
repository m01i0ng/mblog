package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/model"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (u userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := u.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

var _ UserBiz = (*userBiz)(nil)

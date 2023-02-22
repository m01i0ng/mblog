// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package user

import (
	"context"
	"errors"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/m01i0ng/mblog/internal/mblog/store"
	"github.com/m01i0ng/mblog/internal/pkg/errno"
	"github.com/m01i0ng/mblog/internal/pkg/log"
	"github.com/m01i0ng/mblog/internal/pkg/model"
	v1 "github.com/m01i0ng/mblog/pkg/api/mblog/v1"
	"github.com/m01i0ng/mblog/pkg/auth"
	"github.com/m01i0ng/mblog/pkg/token"
	"gorm.io/gorm"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Get(ctx context.Context, username string) (*v1.GetUserResponse, error)
	List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error)
	Update(ctx context.Context, username string, request *v1.UpdateUserRequest) error
	Delete(ctx context.Context, username string) error
}

var _ UserBiz = (*userBiz)(nil)

type userBiz struct {
	ds store.IStore
}

func (b *userBiz) List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error) {
	count, list, err := b.ds.Users().List(ctx, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("Failed to list users from storage", "err", err)
		return nil, err
	}
	users := make([]*v1.UserInfo, 0, len(list))
	for _, item := range list {
		user := item
		users = append(users, &v1.UserInfo{
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	log.C(ctx).Infow("Get users from backend storage", "count", len(users))
	return &v1.ListUserResponse{
		TotalCount: count,
		Users:      users,
	}, nil
}

func (b *userBiz) Update(ctx context.Context, username string, user *v1.UpdateUserRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if user.Email != nil {
		userM.Email = *user.Email
	}
	if user.Phone != nil {
		userM.Phone = *user.Phone
	}
	if user.Nickname != nil {
		userM.Nickname = *user.Nickname
	}

	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

func (b *userBiz) Delete(ctx context.Context, username string) error {
	err := b.ds.Users().Delete(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

func (b *userBiz) Get(ctx context.Context, username string) (*v1.GetUserResponse, error) {
	user, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, user)

	resp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
	return &resp, nil
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'user.username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

func (b *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if err = auth.Compare(userM.Password, r.OldPassword); err != nil {
		return errno.ErrPasswordIncorrect
	}

	userM.Password, err = auth.Encrypt(r.NewPassword)
	if err = b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	if err = auth.Compare(r.Password, user.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: t}, nil
}

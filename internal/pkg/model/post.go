// Copyright 2023 m01i0ng <alley.ma@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/m01i0ng/mblog.

package model

import "time"

type PostM struct {
	Content   string    `gorm:"column:content"`        //
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	ID        int64     `gorm:"column:id;primary_key"` //
	PostID    string    `gorm:"column:postID"`         //
	Title     string    `gorm:"column:title"`          //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
	Username  string    `gorm:"column:username"`       //
}

// TableName sets the insert table name for this struct type
func (p *PostM) TableName() string {
	return "post"
}

package post

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/m01i0ng/mblog/internal/mblog/store"
)

func TestNew(t *testing.T) {
	type args struct {
		ds store.IStore
	}

	tests := []struct {
		name string
		args args
		want *PostController
	}{
		//   TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.ds))
		})
	}
}

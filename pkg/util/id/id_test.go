package id

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenShortID(t *testing.T) {
	shortID := GenShortID()
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortID := GenShortID()
	if shortID == "" {
		b.Error("Failed to gen short id")
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

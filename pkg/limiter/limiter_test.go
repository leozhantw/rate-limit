package limiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	const (
		key = "test-key"

		max = 10
		ttl = 1
	)

	l := New(max, ttl)

	for i := 1; i <= 10; i++ {
		record, err := l.Visit(key)
		assert.NoError(t, err)
		assert.Equal(t, i, record.Count)
		assert.Equal(t, max-i, record.Remaining)
	}
	record, err := l.Visit(key)
	assert.Error(t, err)
	assert.Equal(t, 11, record.Count)
	assert.Equal(t, 0, record.Remaining)

	time.Sleep(1010 * time.Millisecond)

	record, err = l.Visit(key)
	assert.NoError(t, err)
	assert.Equal(t, 1, record.Count)
	assert.Equal(t, 9, record.Remaining)
}

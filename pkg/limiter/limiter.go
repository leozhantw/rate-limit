package limiter

import (
	"errors"
	"sync"
	"time"
)

type Limiter interface {
	Visit(key string) (Record, error)
}

type limiter struct {
	max     int
	ttl     int
	records map[string]*Record
	mux     sync.Mutex
}

func New(max, ttl int) Limiter {
	return &limiter{
		max:     max,
		ttl:     ttl,
		records: make(map[string]*Record),
	}
}

type Record struct {
	Count     int
	Remaining int
	ResetAt   int
}

func (l *limiter) Visit(key string) (Record, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if _, ok := l.records[key]; !ok {
		l.records[key] = &Record{}
	}

	r := l.records[key]
	r.Count++
	if r.ResetAt == 0 {
		now := int(time.Now().Unix())
		r.ResetAt = now + l.ttl
		time.AfterFunc(time.Duration(l.ttl)*time.Second, func() {
			l.mux.Lock()
			delete(l.records, key)
			l.mux.Unlock()
		})
	}
	r.Remaining = l.max - r.Count
	if r.Remaining < 0 {
		r.Remaining = 0
		return *r, errors.New("Limit exceeded")
	}
	return *r, nil
}

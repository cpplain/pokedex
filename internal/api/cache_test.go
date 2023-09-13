package api

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key  string
		data []byte
	}{
		{
			key:  "https://example.com/path1",
			data: []byte("path1 data"),
		},
		{
			key:  "https://example.com/path2",
			data: []byte("path2 data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.add(c.key, c.data)
			data, ok := cache.get(c.key)

			if !ok {
				t.Errorf("expected to find key")
				return
			}

			if string(data) != string(c.data) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCleanCache(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5*time.Millisecond
	const key = "https://example.com"
	cache := NewCache(interval)
	cache.add(key, []byte("data"))

	if _, ok := cache.get(key); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.get(key); ok {
		t.Errorf("expected not to find key")
		return
	}
}

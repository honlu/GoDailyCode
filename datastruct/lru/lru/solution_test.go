package lru

import "testing"

func TestLRUGetMovesNodeAndPreservesEvictionOrder(t *testing.T) {
	cache := NewLRU(2)
	cache.Put("first", 1)
	cache.Put("second", 2)

	if got := cache.Get("first"); got != 1 {
		t.Fatalf("Get(first) = %v, want 1", got)
	}

	cache.Put("third", 3)
	if got := cache.Get("second"); got != nil {
		t.Fatalf("Get(second) = %v, want nil after eviction", got)
	}
	if got := cache.Get("first"); got != 1 {
		t.Fatalf("Get(first) = %v, want 1", got)
	}
	if got := cache.Get("third"); got != 3 {
		t.Fatalf("Get(third) = %v, want 3", got)
	}
}

func TestLRUPutUpdatesExistingValue(t *testing.T) {
	cache := NewLRU(2)
	cache.Put("key", 1)
	cache.Put("key", 2)

	if got := cache.Get("key"); got != 2 {
		t.Fatalf("Get(key) = %v, want 2", got)
	}
	if cache.len != 1 {
		t.Fatalf("cache.len = %d, want 1", cache.len)
	}
}

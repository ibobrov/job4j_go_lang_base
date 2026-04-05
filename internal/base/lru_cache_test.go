package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutAndGetSingleElement(t *testing.T) {
	cache := NewLruCache(2)
	cache.Put("a", "1")

	actual := cache.Get("a")

	assert.Equal(t, "1", *actual)
	assert.Equal(t, "a", cache.Head.Key)
	assert.Equal(t, "a", cache.Tail.Key)
}

func TestGetMissingKeyReturnsNil(t *testing.T) {
	cache := NewLruCache(2)
	cache.Put("a", "1")

	actual := cache.Get("missing")

	assert.Nil(t, actual)
}

func TestPutEvictsLeastRecentlyUsed(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3") // должен вытеснить a

	assert.Nil(t, cache.Get("a"))
	assert.Equal(t, "2", *cache.Get("b"))
	assert.Equal(t, "3", *cache.Get("c"))
}

func TestGetMovesNodeToHead(t *testing.T) {
	cache := NewLruCache(3)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3")
	// порядок: c -> b -> a

	assert.Equal(t, "2", *cache.Get("b"))
	// после Get("b"): b -> c -> a
	assert.Equal(t, "b", cache.Head.Key)
	assert.Equal(t, "a", cache.Tail.Key)
}

func TestPutExistingKeyUpdatesValueAndMovesToHead(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	// b -> a

	cache.Put("a", "10")
	// a -> b

	assert.Equal(t, "10", *cache.Get("a"))
	assert.Equal(t, "a", cache.Head.Key)
	assert.Equal(t, "b", cache.Tail.Key)
}

func TestLRUBehaviorAfterGet(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	// b -> a

	assert.Equal(t, "1", *cache.Get("a"))
	// a -> b

	cache.Put("c", "3")
	// должен вытесниться b

	assert.Nil(t, cache.Get("b"))
	assert.Equal(t, "1", *cache.Get("a"))
	assert.Equal(t, "3", *cache.Get("c"))
}

func TestZeroSizeCache(t *testing.T) {
	cache := NewLruCache(0)

	cache.Put("a", "1")

	assert.Nil(t, cache.Get("a"))
	assert.Nil(t, cache.Head)
	assert.Nil(t, cache.Tail)
}

package base

import "testing"

func TestPutAndGetSingleElement(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")

	got := cache.Get("a")
	if got == nil {
		t.Fatal("expected value, got nil")
	}
	if *got != "1" {
		t.Fatalf("expected 1, got %s", *got)
	}

	if cache.Head == nil || cache.Head.Key != "a" {
		t.Fatal("expected head to be a")
	}
	if cache.Tail == nil || cache.Tail.Key != "a" {
		t.Fatal("expected tail to be a")
	}
}

func TestGetMissingKeyReturnsNil(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")

	got := cache.Get("missing")
	if got != nil {
		t.Fatalf("expected nil, got %v", *got)
	}
}

func TestPutEvictsLeastRecentlyUsed(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3") // должен вытеснить a

	if cache.Get("a") != nil {
		t.Fatal("expected a to be evicted")
	}

	gotB := cache.Get("b")
	if gotB == nil || *gotB != "2" {
		t.Fatal("expected b to be present")
	}

	gotC := cache.Get("c")
	if gotC == nil || *gotC != "3" {
		t.Fatal("expected c to be present")
	}
}

func TestGetMovesNodeToHead(t *testing.T) {
	cache := NewLruCache(3)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3")
	// порядок: c -> b -> a

	got := cache.Get("b")
	if got == nil || *got != "2" {
		t.Fatal("expected value 2 for key b")
	}

	// после Get("b"): b -> c -> a
	if cache.Head == nil || cache.Head.Key != "b" {
		t.Fatalf("expected head to be b, got %v", cache.Head)
	}
	if cache.Tail == nil || cache.Tail.Key != "a" {
		t.Fatalf("expected tail to be a, got %v", cache.Tail)
	}
}

func TestPutExistingKeyUpdatesValueAndMovesToHead(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	// b -> a

	cache.Put("a", "10")
	// a -> b

	got := cache.Get("a")
	if got == nil {
		t.Fatal("expected value, got nil")
	}
	if *got != "10" {
		t.Fatalf("expected updated value 10, got %s", *got)
	}

	if cache.Head == nil || cache.Head.Key != "a" {
		t.Fatal("expected head to be a")
	}
	if cache.Tail == nil || cache.Tail.Key != "b" {
		t.Fatal("expected tail to be b")
	}
}

func TestLRUBehaviorAfterGet(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	// b -> a

	got := cache.Get("a")
	if got == nil || *got != "1" {
		t.Fatal("expected a to be found")
	}
	// a -> b

	cache.Put("c", "3")
	// должен вытесниться b

	if cache.Get("b") != nil {
		t.Fatal("expected b to be evicted")
	}
	if cache.Get("a") == nil {
		t.Fatal("expected a to remain in cache")
	}
	if cache.Get("c") == nil {
		t.Fatal("expected c to remain in cache")
	}
}

func TestZeroSizeCache(t *testing.T) {
	cache := NewLruCache(0)

	cache.Put("a", "1")

	if cache.Get("a") != nil {
		t.Fatal("expected nil for zero-sized cache")
	}
	if cache.Head != nil || cache.Tail != nil {
		t.Fatal("expected empty cache")
	}
}

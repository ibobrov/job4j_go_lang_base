package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	maxSize int
	size    int
	Head    *Node
	Tail    *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		maxSize: size,
	}
}

func (cache *LruCache) Put(key string, value string) {
	if cache.maxSize <= 0 {
		return
	}

	node := cache.findNode(key)
	if node != nil {
		node.Value = value
		cache.moveToHead(node)
		return
	}

	newNode := &Node{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  cache.Head,
	}

	if cache.Head != nil {
		cache.Head.Prev = newNode
	}
	cache.Head = newNode

	if cache.Tail == nil {
		cache.Tail = newNode
	}

	if cache.size < cache.maxSize {
		cache.size++
		return
	}

	cache.removeTail()
}

func (cache *LruCache) Get(key string) *string {
	node := cache.findNode(key)
	if node == nil {
		return nil
	}

	cache.moveToHead(node)
	return &node.Value
}

func (cache *LruCache) findNode(key string) *Node {
	node := cache.Head
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil
}

func (cache *LruCache) moveToHead(node *Node) {
	if node == nil || node == cache.Head {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == cache.Tail {
		cache.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = cache.Head

	if cache.Head != nil {
		cache.Head.Prev = node
	}
	cache.Head = node

	// На случай, если список был пуст
	if cache.Tail == nil {
		cache.Tail = node
	}
}

func (cache *LruCache) removeTail() {
	if cache.Tail == nil {
		return
	}

	if cache.Head == cache.Tail {
		cache.Head = nil
		cache.Tail = nil
		cache.size = 0
		return
	}

	cache.Tail = cache.Tail.Prev
	cache.Tail.Next = nil
}

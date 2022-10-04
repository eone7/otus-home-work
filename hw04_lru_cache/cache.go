package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	/*
		если элемент присутствует в словаре, то обновляем его значение и перемещаем элемент в начало очереди;
		если элемента нет в словаре, то добавляем в словарь и в начало очереди
		(при этом, если размер очереди больше ёмкости кэша,
		то удаляем последний элемент из очереди и его значение из словаря);
		возвращаемое значение - флаг, присутствовал ли элемент в кэше.
	*/
	item := cacheItem{key, value}

	if i, ok := c.items[key]; ok {
		c.items[key].Value = item
		c.queue.MoveToFront(i)
		return true
	}

	if c.queue.Len() >= c.capacity {
		back := c.queue.Back()
		c.queue.Remove(back)
		delete(c.items, back.Value.(cacheItem).key)
	}

	c.queue.PushFront(item)
	c.items[key] = c.queue.Front()
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	/*
		если элемент присутствует в словаре,
		то перемещаем элемент в начало очереди и вовзращааем его значение и true;
		если элемента нет в словаре, то nil и false
	*/
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		item := c.items[key].Value
		return item.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	// Очищаем кэш, удаляем элементы в очереди и из масива значений
	for k, v := range c.items {
		c.queue.Remove(v)
		delete(c.items, k)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

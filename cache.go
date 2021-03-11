package scache

import (
	"container/list"
	"fmt"
)

type Cache struct {
	capacity int
	used     int
	list     *list.List
	table    map[string]*list.Element
}
type entry struct {
	key   string
	value Value
}
type Value interface {
	Len() int
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		used:     0,
		list:     list.New(),
		table:    make(map[string]*list.Element),
	}
}

func (this *Cache) Add(key string, value Value) {
	bytes := value.Len()
	for bytes+this.used >= this.capacity {
		end := this.list.Back()
		entry := end.Value.(*entry)
		key, value := entry.key, entry.value
		this.used += value.Len()
		this.list.Remove(end)
		delete(this.table, key)
	}
	entry := entry{key, value}
	ele := this.list.PushFront(&entry)
	this.used += len(key) + value.Len()
	this.table[key] = ele
}

func (this *Cache) Get(key string) (Value, bool) {
	ele, ok := this.table[key]
	if ok {
		this.list.MoveToFront(ele)
		entry := ele.Value.(*entry)
		return entry.value, true
	} else {
		return nil, false
	}
}
func (this *Cache) Delete(key string) {
	ele, ok := this.table[key]
	if ok {
		this.list.Remove(ele)
		delete(this.table, key)
	}
}

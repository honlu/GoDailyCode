package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
}

type cacheItem struct {
	value   interface{} // 值
	expired time.Time   // 过期时间
}

// Set 指针类型的方法接受者，才可以保证在方法内部访问和修改对象的数据。
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.data.Store(key, cacheItem{
		value:   value,
		expired: time.Now().Add(duration),
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, ok := c.data.Load(key)
	if !ok {
		return nil, false
	}

	cacheItem := item.(cacheItem)
	if time.Now().After(cacheItem.expired) { // 判断是否过期
		c.data.Delete(key)
		return nil, false
	}

	return cacheItem.value, true
}

func main() {
	cache := &Cache{}

	cache.Set("key1", "value1", 5*time.Second)
	cache.Set("key2", "value2", 10*time.Second)

	value1, found1 := cache.Get("key1")
	if found1 {
		fmt.Printf("Value1:%v\n", value1)
	} else {
		fmt.Printf("Key1 not found\n")
	}

	value2, found2 := cache.Get("key2")
	if found2 {
		fmt.Printf("Value2:%v\n", value2)
	} else {
		fmt.Printf("Key2 not found\n")
	}

	time.Sleep(6 * time.Second)

	value1, found1 = cache.Get("key1")
	if found1 {
		fmt.Printf("Value1:%v\n", value1)
	} else {
		fmt.Printf("Key1 not found\n")
	}
	value2, found2 = cache.Get("key2")
	if found2 {
		fmt.Printf("Value2:%v\n", value2)
	} else {
		fmt.Printf("Key2 not found\n")
	}
}

package index

import (
	"fmt"
	"sync"
)

// Basic hash table implementation to use as an index
// hashing is O(n)

type Key interface{}

type Value struct {
	Data interface{}
	Next *Value
}

type HashTable struct {
	Items map[int]Value
	Lock  sync.RWMutex
}

func hash(k Key) (h int) {
	key := fmt.Sprintf("%s", k)

	// TODO find better way to hash ?
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}

	return h
}

func New() *HashTable {
	return &HashTable{
		Items: make(map[int]Value),
	}
}

func (ht *HashTable) Insert(k Key, v interface{}) error {
	// Lock the mutex, so we can implement thread safety for maps

	if ht.Items == nil {
		return fmt.Errorf("Hash table is not initialied, call New")
	}

	// Lock mutex for write
	ht.Lock.Lock()
	defer ht.Lock.Unlock()

	if ht.Items[hash(k)].Data != nil {
		new_elem := &Value{
			Data: v,
			Next: &Value{
				Data: ht.Items[hash(k)].Data,
				Next: ht.Items[hash(k)].Next,
			},
		}
		ht.Items[hash(k)] = *new_elem
	} else {
		new_elem := &Value{}
		new_elem.Data = v
		ht.Items[hash(k)] = *new_elem
	}

	return nil
}

func (ht *HashTable) Search(k Key) (values []interface{}) {
	lookup := ht.Items[hash(k)]

	if lookup.Data != nil {
		values = append(values, lookup.Data)
		for lookup.Next != nil {
			lookup = *lookup.Next
			values = append(values, lookup.Data)
		}
	}

	return values
}

func (ht *HashTable) Length() int {
	return len(ht.Items)
}

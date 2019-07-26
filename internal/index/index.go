package index

import (
	"fmt"
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
	if ht.Items == nil {
		return fmt.Errorf("Hash table is not initialied, call New")
	}

	// we already have something there, let's create a list :D
	// last in, first out
	if len(ht.Search(k)) > 0 {
		new_elem := Value{
			Data: v,
			Next: &Value{
				Data: ht.Items[hash(k)].Data,
				Next: ht.Items[hash(k)].Next,
			},
		}
		ht.Items[hash(k)] = new_elem
	} else {
		ht.Items[hash(k)] = Value{
			Data: v,
		}
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

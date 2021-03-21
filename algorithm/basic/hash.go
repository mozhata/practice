package basic

import (
	"container/list"
	"hash/maphash"
)

const tableCap = 769

type HashSet struct {
	data []list.List
	mph  maphash.Hash
}

func NewHashSet() *HashSet {
	return &HashSet{
		data: make([]list.List, tableCap), //设置一个质数作为底层数组容量
	}
}

func (h *HashSet) hash(key string) uint64 {
	h.mph.Reset()
	h.mph.WriteString(key)
	val := h.mph.Sum64()
	return val % tableCap
}

func (h *HashSet) Add(key string) {
	idx := h.hash(key)
	for e := h.data[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			return // contains
		}
	}
	h.data[idx].PushBack(key)
}

func (h *HashSet) Remove(key string) {
	idx := h.hash(key)
	for e := h.data[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			h.data[idx].Remove(e)
			break
		}
	}
}
func (h *HashSet) Contains(key string) bool {
	idx := h.hash(key)
	for e := h.data[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			return true
		}
	}
	return false
}

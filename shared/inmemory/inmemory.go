package inmemory

import (
	"sync"
	"sync/atomic"
)

// IndexMap InMemoryDBの構造体
type IndexMap struct {
	lock    sync.RWMutex
	counter uint64
	items   map[uint64]interface{}
}

// NewIndexMap IndexMap構造体を生成
func NewIndexMap() *IndexMap {
	return &IndexMap{items: make(map[uint64]interface{})}
}

// Index Counterをプラス1加算
func (m *IndexMap) Index() uint64 {
	return atomic.AddUint64(&m.counter, 1)
}

// Set idx位置に要素を設定
func (m *IndexMap) Set(idx uint64, value interface{}) {
	m.lock.Lock()
	m.items[idx] = value
	m.lock.Unlock()
}

// Get idx位置の要素を取得
func (m *IndexMap) Get(idx uint64) (value interface{}, ok bool) {
	m.lock.RLock()
	v, ok := m.items[idx]
	m.lock.RUnlock()
	return v, ok
}

// Range 結果がfalseになるまで各要素に関数を実行
func (m *IndexMap) Range(f func(idx uint64, value interface{}) bool) {
	m.lock.Lock()
	for k, v := range m.items {
		if !f(k, v) {
			break
		}
	}
	m.lock.Unlock()
}

// Remove idx位置の要素を削除
func (m *IndexMap) Remove(idx uint64) {
	m.lock.Lock()
	delete(m.items, idx)
	m.lock.Unlock()
}

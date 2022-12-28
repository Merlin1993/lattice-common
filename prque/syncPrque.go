package prque

import (
	"sync"
)

type SyncPrQue struct {
	*PrQue
	lock sync.Mutex
}

func NewSyncPrQue(setIndex SetIndexCallback) *SyncPrQue {
	return &SyncPrQue{PrQue: New(setIndex)}
}

func (p *SyncPrQue) Push(data interface{}, priority int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.PrQue.Push(data, priority)
}

func (p *SyncPrQue) Pop() (interface{}, int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.PrQue.Pop()
}

func (p *SyncPrQue) PopItem() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.PrQue.PopItem()
}

func (p *SyncPrQue) Remove(i int) interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.PrQue.Remove(i)
}

package prque

import (
	"sync"
)

//具有优先级的集合
type PrQueSet struct {
	que *PrQue
	set map[interface{}]struct{}

	lock sync.RWMutex
}

func NewPrQueSet() *PrQueSet {
	return &PrQueSet{
		que: New(nil),
		set: make(map[interface{}]struct{}),
	}
}

func (p *PrQueSet) Push(data interface{}, priority int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, ok := p.set[data]; !ok {
		p.que.Push(data, priority)
		p.set[data] = struct{}{}
	}
}

func (p *PrQueSet) Peek() (data interface{}, priority int64) {
	return p.que.Peek()
}

func (p *PrQueSet) Pop() (interface{}, int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	value, pro := p.que.Pop()
	delete(p.set, value)
	return value, pro
}

func (p *PrQueSet) Size() int {
	return p.que.Size()
}

func (p *PrQueSet) Reset() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.que.Reset()
	p.set = make(map[interface{}]struct{})
}

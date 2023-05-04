package lattice_common

import "sync"

type FIFOSet struct {
	queue []interface{}
	set   map[interface{}]struct{}

	lock sync.RWMutex
}

func NewFIFOSet() *FIFOSet {
	return &FIFOSet{
		queue: make([]interface{}, 0),
		set:   make(map[interface{}]struct{}),
	}
}

func (p *FIFOSet) Push(data interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, ok := p.set[data]; !ok {
		p.queue = append(p.queue, data)
		p.set[data] = struct{}{}
	}
}

func (p *FIFOSet) PushAll(dataSlice []interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, data := range dataSlice {
		if _, ok := p.set[data]; !ok {
			p.queue = append(p.queue, data)
			p.set[data] = struct{}{}
		}
	}
}

func (p *FIFOSet) Peek() (data interface{}) {
	if len(p.queue) == 0 {
		return nil
	}
	return p.queue[0]
}

func (p *FIFOSet) PeekAll() []interface{} {
	return p.queue
}

func (p *FIFOSet) Pop(size int) []interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	if len(p.queue) == 0 {
		return nil
	}
	if size > len(p.queue) {
		size = len(p.queue)
	}
	value := p.queue[:size-1]
	p.queue = p.queue[size:]
	for _, item := range value {
		delete(p.set, item)
	}
	return value
}

func (p *FIFOSet) PopAll() []interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	value := p.queue
	p.reset()
	return value
}

func (p *FIFOSet) Size() int {
	return len(p.queue)
}

func (p *FIFOSet) Reset() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.reset()
}

func (p *FIFOSet) reset() {
	p.queue = make([]interface{}, 0)
	p.set = make(map[interface{}]struct{})
}

func (p *FIFOSet) Purge(data interface{}) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	var isExist = false
	for i := 0; i < len(p.queue); {
		if data == p.queue[i] {
			p.queue = append(p.queue[:i], p.queue[i+1:]...)
			isExist = true
		} else {
			i++
		}
	}
	return isExist
}

//遍历删除有问题的
func (p *FIFOSet) PurgeAll(data []interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	dataMap := make(map[interface{}]struct{})
	for _, item := range data {
		dataMap[item] = struct{}{}
	}
	for i := 0; i < len(p.queue); {
		if _, ok := dataMap[p.queue[i]]; ok {
			p.queue = append(p.queue[:i], p.queue[i+1:]...)
		} else {
			i++
		}
	}
}

func (p *FIFOSet) Copy() *FIFOSet {
	cpy := NewFIFOSet()
	cpy.PushAll(p.PeekAll())
	return cpy
}

//首先可以取出一部分,而且是有序的
//其次可以根据上次取出来的部分，继续取部分（或者不用？直接用广播打出去？）
//最后，支持与一列表进行对比搜索，删除所有已经成功的
//

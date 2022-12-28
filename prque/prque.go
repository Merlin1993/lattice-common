package prque

import "container/heap"

type PrQue struct {
	cont *sortStack
}

func New(setIndex SetIndexCallback) *PrQue {
	return &PrQue{newSortStack(setIndex)}
}

func (p *PrQue) Push(data interface{}, priority int64) {
	heap.Push(p.cont, &item{data, priority})
}

func (p *PrQue) Peek() (interface{}, int64) {
	item := p.cont.blocks[0][0]
	return item.value, item.priority
}

func (p *PrQue) Pop() (interface{}, int64) {
	temp := heap.Pop(p.cont)
	if temp == nil {
		return nil, 0
	}
	item := temp.(*item)
	return item.value, item.priority
}

func (p *PrQue) PopItem() interface{} {
	return heap.Pop(p.cont).(*item).value
}

func (p *PrQue) Remove(i int) interface{} {
	if i < 0 {
		return nil
	}
	return heap.Remove(p.cont, i)
}

func (p *PrQue) Empty() bool {
	return p.cont.Len() == 0
}

func (p *PrQue) Size() int {
	return p.cont.Len()
}

func (p *PrQue) Reset() {
	*p = *New(p.cont.setIndex)
}

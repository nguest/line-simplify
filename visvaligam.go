package main

import (
	"math"
)

type Item struct {
	a    float64
	pIdx int
	next *Item
	prev *Item
	idx  int
}

func Visvalingam(data []Datum, count int) []Datum {
	var idxMap []int
	if len(data) <= count {
		idxMap = make([]int, len(data))
		for i := range data {
			idxMap[i] = i
		}
	}

	// threshoold?

	removed := 0

	// build initial heap
	heap := minHeap(make([]*Item, 0, len(data)))

	linkedListStart := &Item{
		a:    math.Inf(1),
		pIdx: 0,
	}
	heap.Push(linkedListStart)

	// make path Items, exclude start and end
	items := make([]Item, len(data))

	prev := linkedListStart
	for i := 1; i < len(data)-1; i++ {
		item := &items[i]
		item.a = area(data, i-1, i, i+1)
		item.pIdx = i
		item.prev = prev

		heap.Push(item)
		prev.next = item
		prev = item
	}

	// make final Item
	lastItem := &items[len(data)-1]
	lastItem.a = math.Inf(1)
	lastItem.pIdx = len(data) - 1
	lastItem.prev = prev
	prev.next = lastItem
	heap.Push(lastItem)

	// now the points reducer

	for len(heap) > 0 {
		curr := heap.Pop()
		if len(data)-removed <= count {
			break
		}
		next := curr.next
		prev := curr.prev

		// remove item from list
		prev.next = curr.next
		next.prev = curr.prev
		removed++

		// calculate new areas
		if prev.prev != nil {
			a := area(data, prev.prev.pIdx, prev.pIdx, next.pIdx)
			a = math.Max(a, curr.a)
			heap.Update(prev, a)
		}

		if next.next != nil {
			a := area(data, prev.pIdx, next.pIdx, next.next.pIdx)
			a = math.Max(a, curr.a)
			heap.Update(next, a)
		}
	}

	item := linkedListStart

	cnt := 0
	for item != nil {
		data[cnt] = data[item.pIdx]
		cnt++
		item = item.next
	}
	return data[:cnt]
}

type minHeap []*Item

func (h *minHeap) Push(item *Item) {
	item.idx = len(*h)
	*h = append(*h, item)
	h.up(item.idx)
}

func (h *minHeap) Pop() *Item {
	removed := (*h)[0]
	lastItem := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]

	if len(*h) > 0 {
		lastItem.idx = 0
		(*h)[0] = lastItem
		h.down(0)
	}

	return removed
}

func (h *minHeap) Update(item *Item, a float64) {
	if item.a > a {
		// area became smaller
		item.a = a
		h.up(item.idx)
	} else {
		// area became larger
		item.a = a
		h.down(item.idx)
	}
}

func (h minHeap) up(i int) {
	obj := h[i]
	for i > 0 {
		up := ((i + 1) >> 1) - 1
		parent := h[up]
		if parent.a <= obj.a {
			// parent smaller so get out of heap ops
			break
		}

		// swap nodes around
		parent.idx = i
		h[i] = parent
		obj.idx = up
		h[up] = obj
	}
}

func (h minHeap) down(i int) {
	obj := h[i]

	for {
		right := (1 + i) << 1
		left := right - 1
		down := i
		child := h[down]

		// swap with smallest child
		if left < len(h) && h[left].a < child.a {
			down = left
			child = h[down]
		}

		if right < len(h) && h[right].a < child.a {
			down = right
			child = h[down]
		}

		// quit if none smaller
		if down == i {
			break
		}

		// swap nodes around
		child.idx = i
		h[child.idx] = child
		obj.idx = down
		h[down] = obj
		i = down
	}
}

// returns double the triangle area
func area(data []Datum, i0, i1, i2 int) float64 {
	v0 := data[i0]
	v1 := data[i1]
	v2 := data[i2]

	return math.Abs(
		v0.Lon*(v1.Lat-v2.Lat) + v1.Lon*(v2.Lat-v0.Lat) + v2.Lon*(v0.Lat-v1.Lat))
}

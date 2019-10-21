package main

import "math"

// type TriP struct {
// 	v0, v1, v2 Datum
// 	a          float64
// }
// type Triangle struct {
// 	p    TriP
// 	prev *Triangle
// 	next *Triangle
// }

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
	heap := minHeap(make([]*Item, len(data)))

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
}

func (h *minHeap) Pop() *Item {
	item.idx = len(*h)
}

func (h *minHeap) Update(item *Item, a float64) {
	item.idx = len(*h)
}

// 	heap := []Triangle{} //minHeap()
// 	maxA := 0.0
// 	t := Triangle{}
// 	ts := []Triangle{}

// 	//data = data.map(function (d) { return d.slice(0,2); });
// 	for i, v := range data {
// 		data[i] = v
// 	}

// 	// calculate triangles and their areas
// 	for i := 1; i < len(data)-2; i++ {
// 		t := Triangle{
// 			p: TriP{
// 				v0: data[i-1],
// 				v1: data[i],
// 				v2: data[i+1],
// 				a:  area(t),
// 			},
// 		}

// 		if t.p.a > 0 {
// 			ts = append(ts, t)
// 			Push(heap, t)
// 		}
// 	}

// 	for i := 0; i < len(ts)-1; i++ {
// 		t = ts[i]
// 		t.prev = &ts[i-1]
// 		t.next = &ts[i+1]
// 	}
// 	///////
// 	if t.p.a < maxA {
// 		t.p.a = maxA
// 	} else {
// 		maxA = t.p.a
// 	}

// 	if t.prev != nil {
// 		t.prev.next = t.next
// 		t.prev.p[2] = t.p[2]
// 		update(heap, t.prev)
// 	} else {
// 		t.p[0] = t.a
// 	}

// }
// returns double the triangle area
func area(data []Datum, i0, i1, i2 int) float64 {
	v0 := data[i0]
	v1 := data[i1]
	v2 := data[i2]

	return math.Abs(
		v0.Lon*(v1.Lat-v2.Lat) + v1.Lon*(v2.Lat-v0.Lat) + v2.Lon*(v0.Lat-v1.Lat))
}

// func minHeap() {

// }

// func Push(heap []Triangle, t Triangle) int {
// 	//for i := 0, n = arguments.length; i < n; ++i {
// 	var object = t
// 	heap = append(heap, object)
// 	Up(heap, append(heap, object))

// 	return len(heap)
// }

// func Up(ts []Triangle, i int) {
// 	object := ts[i]
// 	for i > 0 {
// 		up := ((i + 1) >> 1) - 1
// 		parent := ts[up]
// 		if CompareArea(object, parent) >= 0 {
// 			break
// 		}
// 		ts[i] = parent
// 		ts[up] = object
// 	}
// }

// func Down(ts []Triangle, i int) {
// 	var object = ts[i]
// 	right := (i + 1) * 2
// 	left := right - 1
// 	down := i
// 	child := ts[down]
// 	for {
// 		if left < len(ts) && CompareArea(ts[left], child) < 0 {
// 			child = ts[left]
// 		}
// 		if right < len(ts) && CompareArea(ts[right], child) < 0 {
// 			child = ts[right]
// 		}
// 		if down == i {
// 			break
// 		}
// 		ts[i] = child
// 		ts[down] = object
// 	}
// }

// func CompareArea(t1, t2 Triangle) float64 {
// 	return t1.a - t2.a
// }

package lrucache

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i int, j int) bool {
	return h[i].time.Before(h[j].time)
}

func (h Heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

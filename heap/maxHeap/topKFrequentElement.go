package maxHeap

func topKFrequent(nums []int, nn int) []int {
	m := make(map[int]int)
	for _, t := range nums {
		if cnt, ok := m[t]; ok {
			m[t] = cnt + 1
		} else {
			m[t] = 1
		}
	}
	h := make(heap, 0)
	for k, v := range m {
		n := node{
			count: v,
			num:   k,
		}
		h.insert(n)
	}
	res := []int{}
	for i := 0; i < nn; i++ {
		t := h.remove()

		res = append(res, t.num)
	}
	return res
}

type node struct {
	num, count int
}

type heap []node

func (h *heap) insert(q node) {
	(*h) = append((*h), q)
	li := len(*h) - 1
	pr := parent(li)
	for pr >= 0 && (*h)[li].count > (*h)[pr].count {
		(*h)[li], (*h)[pr] = (*h)[pr], (*h)[li]
		li = pr
		pr = parent(li)
	}
}

func (h *heap) remove() node {
	t := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	(*h) = (*h)[:len((*h))-1]
	h.maxHeapify(0)
	return t
}

func (h *heap) maxHeapify(i int) {
	max := i
	l := lc(i)
	if l >= len(*h) {
		return
	}
	r := rc(i)
	if r >= len(*h) {
		max = l
	} else {
		if (*h)[r].count < (*h)[l].count {
			max = l
		} else {
			max = r
		}
	}
	if (*h)[max].count > (*h)[i].count {
		(*h)[max], (*h)[i] = (*h)[i], (*h)[max]
		h.maxHeapify(max)
	} else {
		return
	}

}

func lc(n int) int {
	return n*2 + 1
}

func rc(n int) int {
	return n*2 + 2
}

func parent(n int) int {
	return (n - 1) / 2
}

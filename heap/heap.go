package heap

import "fmt"

type Heap struct {
	arr  []int
	size int
}

func NewHeap() *Heap {
  return &Heap{
    arr: make([]int, 1),
    size: 0,
  }
}

func (h *Heap) Push(value int) {
	h.size++
  if h.size >= len(h.arr) {
    h.resize()
  }
	h.arr[h.size] = value
	
  for i := h.size; i > 1; i >>= 1 {
		if h.arr[i] > h.arr[i/2] {
			h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		} else {
			break
		}
	}
}

func (h *Heap) resize() {
  newArr := make([]int, len(h.arr)*2)
  copy(newArr, h.arr)
  h.arr = newArr
}

func (h *Heap) Pop() int {
  maxItem := h.Peak()
  h.arr[1] = h.arr[h.size]
  h.size--
	
  for i := 1; i*2 <= h.size; {
		ch := i * 2
		if ch+1 <= h.size && h.arr[ch] < h.arr[ch+1] {
			ch++
		}
		if h.arr[ch] > h.arr[i] {
			h.arr[ch], h.arr[i] = h.arr[i], h.arr[ch]
		} else {
			break
		}
    i = ch
	}
	return maxItem
}

func (h *Heap) Peak() int {
  return h.arr[1]
}

/*
func main() {
	h := &Heap{
		arr:  make([]int, 10),
		size: 0,
	}
	h.Push(9)
	h.Push(8)
	h.Push(5)
	h.Push(7)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(10)

	fmt.Println(h)

  fmt.Println(h.Pop())
	fmt.Println(h)
}
*/

package main

import (
	"container/heap"
	//"fmt"
)

type File struct {
	path string
	size int64
}

type FileHeap struct {
	maxRecords int
	files      []File
}

func (h FileHeap) Len() int           { return len(h.files) }
func (h FileHeap) Less(i, j int) bool { return h.files[i].size < h.files[j].size }
func (h FileHeap) Swap(i, j int)      { h.files[i], h.files[j] = h.files[j], h.files[i] }

func (h *FileHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.files = append(h.files, x.(File))
}

func (h *FileHeap) Pop() interface{} {
	old := h.files
	n := len(old)
	x := old[n-1]
	h.files = old[0 : n-1]
	return x
}

func (fh *FileHeap) Check() {
	if len(fh.files) > fh.maxRecords {
		heap.Pop(fh)
	}
}

func (fh *FileHeap) Init() {
	heap.Init(fh)
}

func (fh *FileHeap) Popx() interface{} {
	return heap.Pop(fh)
}

func (fh *FileHeap) Pushx(x interface{}) {
	heap.Push(fh, x)
	fh.Check()
}

// // This example inserts several ints into an IntHeap, checks the minimum,
// // and removes them in order of priority.
// func main() {
// 	fh := &FileHeap{
// 		maxRecords: 3,
// 		files: []File{
// 			{"/", 100},
// 			{"/dir2", 200},
// 			{"/dir6", 600},
// 		},
// 	}
// 	file4 := File{
// 		"/dir4", 400,
// 	}
// 	file5 := File{
// 		"/dir5", 500,
// 	}
// 	file7 := File{
// 		"/dir7", 700,
// 	}
// 	// heap.Init(fh)
// 	// fh.Push(file4)
// 	// fh.Check()
// 	// fh.Push(file5)
// 	// fh.Check()
// 	// if len(fh.files) > fh.maxRecords {
// 	// 	heap.Pop(fh)
// 	// }
// 	// for fh.Len() > 0 {
// 	// 	fmt.Println(heap.Pop(fh).(File).path)
// 	// }
// 	fh.Init()
// 	fh.Pushx(file4)
// 	fh.Pushx(file5)
// 	fh.Pushx(file7)
// 	for fh.Len() > 0 {
// 		fmt.Println(fh.Popx().(File).path)
// 	}
// }

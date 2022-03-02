package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var result = make(map[int]*FileHeap)

func checkPath(path string, cur, depth, maxFile int, result map[int]*FileHeap) map[int]*FileHeap {
	// 初始化文件堆栈
	if cur < depth || depth == 0 {
		if result[cur] == nil {
			files := []File{}
			result[cur] = &FileHeap{
				maxRecords: maxFile,
				files:      files,
			}
		}

		rd, _ := ioutil.ReadDir(path)
		for _, fi := range rd {
			if fi.IsDir() {
				fiPath := fmt.Sprintf("%s/%s", path, fi.Name())
				file := File{
					path: fiPath,
					size: fi.Size(),
				}
				result[cur].Pushx(file)
				// 递归执行遍历所有目录文件
				checkPath(fiPath, cur+1, depth, maxFile, result)
			}

		}
	}
	return result
}

func main() {
	var path string
	var depth int
	var maxFile int

	flag.StringVar(&path, "p", "./", "目录路径, 默认为当前路径")
	flag.IntVar(&depth, "d", 0, "搜索深度, 默认为最深层级搜索")
	flag.IntVar(&maxFile, "m", 10, "输出最大目录条目, 默认为10")
	flag.Parse()

	result = checkPath(path, 0, depth, maxFile, result)
	for i := 0; i < len(result); i++ {
		if result[i] != nil && len(result[i].files) != 0 {
			fmt.Printf("输出第%d级最大%d的目录列表\n", i+1, maxFile)
			filsNum := len(result[i].files)
			for j := 0; j < filsNum; j++ {
				file := result[i].Popx().(File)
				fmt.Printf("%s, %d\n", file.path, file.size)
			}
		}
	}
}

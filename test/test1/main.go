package main

import (
	"fmt"
	"runtime"
	"sync"
)

type MySlice [][]int

func (s *MySlice) Less(i, j int) bool {
	return (*s)[i][0] < (*s)[j][0]
}

func (s *MySlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]

}

func (s *MySlice) Len() int {
	return len(*s)
}

func Print() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(1)
	// res := make([][]int, 0, 10)
	// tmpRes := MySlice(res)nums[i]
	// tmp := []int{1, 3}
	// tmp1 := []int{0, 4}
	// tmp2 := []int{5, 7}
	// tmpRes = append(tmpRes, tmp2)
	// tmpRes = append(tmpRes, tmp)
	// tmpRes = append(tmpRes, tmp1)
	// sort.Sort(&tmpRes)
	// fmt.Println(tmpRes)
	// Print()
	// time.Sleep(time.Second)
}

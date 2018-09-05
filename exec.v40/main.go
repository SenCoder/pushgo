package main

import (
	"sort"
	"fmt"
	"runtime/pprof"
	"flag"
	"os"
	"runtime"
	"log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")




func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	return ret
	sort.Ints(nums)
	length := len(nums)
	for i, v := range nums {
		start := 0
		end := length -1
		for start < end {
			if nums[start] + nums[end] + v > 0 {
				end --
			} else if nums[start] + nums[end] + v < 0 {
				start ++
			} else {
				if i != start && i != end {
					ret = append(ret, []int{nums[i], nums[start], nums[end]})
					fmt.Println(ret)
					break
				}
			}
		}
	}
	return ret
}

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))


	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}
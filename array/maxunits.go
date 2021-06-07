package array

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	fmt.Println(boxTypes)
	var res int
	for _, v := range boxTypes {
		if truckSize <= v[0] {
			res += v[1] * truckSize
			break
		} else {
			truckSize -= v[0]
			res += v[0] * v[1]
		}
	}
	return res

}

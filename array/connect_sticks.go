package array

import "sort"

func connectSticks(sticks []int) int {
	if len(sticks) == 0 {
		return 0
	}
	sort.Ints(sticks)
	res := 0
	for len(sticks) > 1 {
		//connect shortest 2 sticks always
		newStick := sticks[0] + sticks[1]
		res += newStick
		sticks = sticks[2:]
		//add the connected stick back to the sorted sticks list
		pos := insert(sticks, newStick)
		sticks = append(sticks, 0)
		copy(sticks[pos+1:], sticks[pos:])
		sticks[pos] = newStick
	}
	return res
}

//binary search, get the right position to insert the connected stick
func insert(sticks []int, n int) int {
	if len(sticks) == 0 {
		return 0
	}
	l := 0
	r := len(sticks) - 1

	for l+1 < r {
		mid := (l + r) / 2
		if sticks[mid] == n {
			return mid
		}
		if sticks[mid] > n {
			r = mid
		} else {
			l = mid
		}
	}
	if sticks[l] >= n {
		return l
	}
	if sticks[r] >= n {
		return r
	}
	return r + 1
}

package array

func twoSum(nums []int, target int) []int {
	// create a map that contains m[nums[i] = i
	// if map contains nums such that num = target - nums[i], return, map[num] and i

	// nums = [2,7,11,15], target = 9
	// map[2:0, 7:1, 11:2, 15:3]
	// comp = (target - num[i]): 7, 2, -4, -6
	// return [0, 1]
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		val, ok := m1[comp]
		if ok {
			return []int{val, i}
		}
		m1[nums[i]] = i
	}

	return nil
}

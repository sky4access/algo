package array

// {1,2,3} -> {1,2,4}
func addOne(arr []int) []int {
	carry := 1
	result := make([]int, len(arr))
	for i := len(arr) -1; i>= 0; i-- {
		sum := arr[i] + carry
		if sum == 10 {
			carry = 1
		} else {
			carry = 0
		}
		result[i] = sum % 10
	}

	if carry == 1 {
		result = append([]int{1}, result...)
	}
	return result
}

func addNumber(arr []int, num int) []int {
	carry := num
	result := make([]int, len(arr))
	for i := len(arr) -1; i>= 0; i-- {
		sum := arr[i] + carry
		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}
		result[i] = sum % 10
	}

	if carry == 1 {
		result = append([]int{1}, result...)
	}
	return result
}
package array

// hasCommon returns true if two arrays contain common item
func hasCommon(a []int, b []int) bool {

	if len(a)==0 || len(b)==0 {
		return false
	}
	m := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		m[a[i]]=true
	}
	for i :=0; i < len(b); i++ {
		if m[b[i]] {
			return true
		}
	}

	// O (a + b) time, O (a) space
	return false
}


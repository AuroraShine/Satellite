package sorts

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		k := false
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				k = true
			}
		}
		if !k {
			return
		}
	}
}

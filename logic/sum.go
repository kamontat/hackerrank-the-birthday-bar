package logic

func Sum(s []int32, d int32, m int32) (count int32) {
	var sum []int32 = make([]int32, 102)
	sum[0] = 0

	for i := range s {
		sum[i+1] = sum[i] + s[i]
	}

	for i := 0; i < len(s)-int(m); i++ {
		if sum[i+int(m)]-sum[i] == d {
			count++
		}
	}

	return count
}

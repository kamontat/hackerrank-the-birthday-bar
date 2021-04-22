package logic

func sum(s []int32) (a int32) {
	for _, elem := range s {
		a += elem
	}
	return
}

func Normal(s []int32, d int32, m int32) int32 {
	var count int32
	var size int = len(s)

	for i := range s {
		last := i + int(m)

		if last > size {
			return count
		}

		batch := s[i:last]
		if sum(batch) == d {
			count++
		}
	}

	return count
}

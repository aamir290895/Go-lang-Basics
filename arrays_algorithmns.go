package main

func arrangeLowestNextToHighest(arr []int) []int {

	// arr := []int {1,2,25,8,7}

	s := 25
	h := 0
	hi := 0
	si := 0

	for i, v := range arr {

		if v > h {

			h = v
			hi = i
		}

		if s > v {
			s = v
			si = i

		}
	}

	// temp := []int{}

	// for _, v := range arr {

	// 	if v == h && v != s {
	// 		temp = append(temp, v)
	// 		temp = append(temp, s)

	// 	} else if v != s {
	// 		temp = append(temp, v)

	// 	}
	// }

	arr = append(arr[:si], arr[si+1:]...)

	temp := []int{}

	temp = append(temp, arr[s+1:]...)
	arr = append(arr[:hi], s)

	arr = append(arr, temp...)

	return arr
}

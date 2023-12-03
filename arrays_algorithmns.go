package main

func arrangeLowestNextToHighest(arr []int) []int {

	s := 25
	h := 0

	for _, v := range arr {

		if v > h {

			h = v
		}

		if s > v {
			s = v

		}
	}

	temp := []int{}

	for _, v := range arr {

		if v == h && v != s {
			temp = append(temp, v)
			temp = append(temp, s)

		} else if v != s {
			temp = append(temp, v)

		}
	}

	return temp
}

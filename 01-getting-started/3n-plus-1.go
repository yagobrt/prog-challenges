package gettingstarted

import (
	"fmt"
)

func calculate_length(n int, lenghts map[int]int) int {
	_, ok := lenghts[n]
	if !ok {
		var next int
		if n%2 == 0 {
			next = n / 2
		} else {
			next = 3*n + 1
		}
		lenghts[n] = 1 + calculate_length(next, lenghts)
	}
	return lenghts[n]
}

func ThreeNPlusOne() {
	cycle_lengths := make(map[int]int)
	cycle_lengths[1] = 1

	var i, j int

	for {
		_, err := fmt.Scanf("%d %d\n", &i, &j)
		if err != nil {
			return
		}

		start := min(i, j)
		end := max(i, j)

		max_len := 0
		var l int
		for num := start; num <= end; num++ {
			l = calculate_length(num, cycle_lengths)
			max_len = max(l, max_len)
		}
		fmt.Printf("%d %d %d\n", i, j, max_len)
	}
}

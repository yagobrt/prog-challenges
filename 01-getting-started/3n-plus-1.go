package main

import (
	"fmt"
)

func calcular_longitud(n int, longitudes map[int]int) int {
	_, ok := longitudes[n]
	if !ok {
		var sig int
		if n%2 == 0 {
			sig = n / 2
		} else {
			sig = 3*n + 1
		}
		longitudes[n] = 1 + calcular_longitud(sig, longitudes)
	}
	return longitudes[n]
}

func main() {
	long_ciclos := make(map[int]int)
	long_ciclos[1] = 1

	var i, j int

	for {
		_, err := fmt.Scanf("%d %d\n", &i, &j)
		if err != nil {
			return
		}

		inicio := min(i, j)
		final := max(i, j)

		long_max := 0
		var l int
		for num := inicio; num <= final; num++ {
			l = calcular_longitud(num, long_ciclos)
			long_max = max(l, long_max)
		}
		fmt.Printf("%d %d %d\n", i, j, long_max)
	}

}

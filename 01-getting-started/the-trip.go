package gettingstarted

import "fmt"

func TheTrip() {
	var numStudets int
	for {
		// read input
		_, err := fmt.Scanf("%d", &numStudets)
		if err != nil {
			return
		}
		if numStudets == 0 {
			return
		}

		payments := make([]int, numStudets)
		totalAmount := 0
		var dollars int
		var cents int
		for i := 0; i < numStudets; i++ {
			_, err := fmt.Scanf("%d.%d", &dollars, &cents)
			if err != nil {
				return
			}
			payments[i] = dollars*100 + cents
			totalAmount += dollars*100 + cents
		}

		mean := totalAmount / numStudets
		overMean := totalAmount % numStudets
		numChanges := 0
		change := 0
		for _, v := range payments {
			if v > mean {
				change += v - mean
				numChanges += 1
			}
		}

		// distribute back the change to the students who paid more than the mean
		change -= min(numChanges, overMean)

		// Format
		fmt.Printf("$%.2f\n", float32(change)/100)
	}
}

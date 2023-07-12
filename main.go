package main

import (
	"fmt"
	"math"
)

func main() {
	start := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 3}
	for {
		if calculate(start) == 200 {
			fmt.Println(printRes(start))
		}
		if ok := next(start); !ok {
			break
		}
	}
}

// Creating next combination, return value shows whether the operation is successful
func next(input []int) bool {
	i := 0
	needToUpdate := true
	for needToUpdate {
		switch input[i] {
		case 0, 1:
			input[i]++
			needToUpdate = false
		case 2:
			input[i] = 0
			i++
		case 3:
			return false
		}
	}
	return true
}

// Calculating sum of combination, return sum
func calculate(input []int) int64 {
	var res int64 = 0
	var tmp int64 = 0
	j := 0
	for i := 0; i < 10; i++ {
		tmp = tmp + int64(i*int(math.Pow10(j)))
		val := input[i]
		switch val {
		case 0:
			j++
		case 1:
			res -= tmp
			tmp = 0
			j = 0
		case 2:
			res += tmp
			tmp = 0
			j = 0
		case 3:
			res += tmp
		}
	}
	return res
}

// Transforming slice of good combination to result string, return result string
func printRes(input []int) string {
	res := "0"
	for i := 1; i < 10; i++ {
		val := input[i-1]
		switch val {
		case 0, 3:
			res = fmt.Sprintf("%d%s", i, res)
		case 1:
			res = fmt.Sprintf("%d-%s", i, res)
		case 2:
			res = fmt.Sprintf("%d+%s", i, res)
		}
	}
	return res
}

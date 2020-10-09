package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	numbers := make([]int, 2)
	numbers[0] = 0
	numbers[1] = 1
	count := 0

	return func() int {
		if count > 1 {
			numbers = append(numbers, numbers[count-2]+numbers[count-1])
		}

		ret := numbers[count]
		count += 1
		return ret

	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

import "fmt"

func lemonadeChange(bills []int) bool {

	five := 0
	ten := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {

			if five == 0 {
				return false
			} else {
				five--;
				ten ++;
			}

		} else if bills[i] == 20 {
			if ten >= 1 && five >= 1 {
				ten--;
				five--;
			} else if five >= 3 {
				five -= 3;
			}else
			{
				return false
			}
		}
	}
	return true
}

func main() {
	bills := []int{5,5,5,10,20}
	output := lemonadeChange(bills)
	fmt.Println(output)
}

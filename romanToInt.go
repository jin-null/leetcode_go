package main

import "fmt"

func romanToInt(str string) int {
	res:=0
	m:=map[byte] int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(str) - 1; i >= 0; i-- {
		temp := m[str[i]]

		sign := 1
		if temp < last {
			//小数在大数的左边，要减去小数
			sign = -1
		}

		res += sign * temp

		last = temp
	}

	return res
}

func main() {
	str:="MCMXCIV"
	output := romanToInt(str)
	fmt.Println(output)
}

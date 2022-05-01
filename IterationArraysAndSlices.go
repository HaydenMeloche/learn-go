package main

func Looping(character string, amount int) string {
	var repeated string
	for i := 0; i < amount; i++ {
		repeated = repeated + character
	}
	return repeated
}

func sumOfArray(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // after excute
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func sayTotalNumber(total int) {
	fmt.Println("total is ", total)
}
func superAdd(numbers ...int) int {

	total := 0
	//for i := 0 ; i < len(numbers) ; i++
	for _, number := range numbers {
		total += number
	}
	defer sayTotalNumber(total)
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func main() {

	// total, upper := lenAndUpper("jgi")
	// fmt.Println(total, upper)
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("-->", total)

	fmt.Println(canIDrink(2))

}

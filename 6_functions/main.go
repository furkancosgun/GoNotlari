package main

import "fmt"

// Geri Donuş Degeri Olmayan Fonksiyonlar
func sum(x int, y int) {
	fmt.Println(x + y)
}

// Geri Donuş Degeri Olan Fonksiyonlar
func sumAndGet(x int, y int) int {
	return x + y
}

// 2 parametre alır
func calculate(x int, y int) (int, int) {
	return x + y, x - y
}

// Liste Alır
func sum_with_list(int_list []int) int {
	var result int

	for _, value := range int_list {
		result += value
	}
	return result
}

// Sonsuz Arguman Alabilir
func sum_all(arr ...int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}

func main() {
	sum(10, 20)

	total := sumAndGet(20, 1000)
	fmt.Println(total)

	total, diff := calculate(10, 10)

	fmt.Printf("Total:%v ,Diff:%v\n", total, diff)

	list_result := sum_with_list([]int{10, 20, 30, 40, 50, 60, 70, 80, 90})
	fmt.Printf("Result:%v \n", list_result)

	fmt.Println(sum_all(1, 2, 3, 4, 5, 6, 7, 8, 9, 0123, 3*12, 3123*3))
}

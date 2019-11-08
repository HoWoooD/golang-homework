package main

import (
	"fmt"
	"strings"
	"strconv"
	)

func main()  {
	//1
	fmt.Println("First task")
	fmt.Println(multiply(1.12425, 3.5345245))
	fmt.Println(multiply(4, 1.5))
	fmt.Println(multiply(100.5, 2.5))
	fmt.Println(multiply(7.7, 3.1))
	fmt.Println(multiply(3.14, 3.14))

	//2
	fmt.Println("\nSecond task")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacciLoop(i), " ")
	}

	//3
	fmt.Println("\n\nThird task")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacciRecursion(i), " ")
	}

	//4
	fmt.Println("\n\nFourth task")
	arr := []float32{2,5,1,4,3}
	bubbleSort(&arr)
	fmt.Println(arr)

	//5
	fmt.Println("\nFifth task")
	// arr1 := []int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	// arr1 := []int{1, 1, 1, 1, 1, 2, 2, 3, 2}
	arr1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(uniqueCount(&arr1))
}

func multiply(x float64, y float64) float64  {
	//get len after dot
	arr1 :=strings.Split(fmt.Sprintf("%g", x), ".")
	arr2 :=strings.Split(fmt.Sprintf("%g", y), ".")
	lenAfterComX := 0
	lenAfterComY := 0
	if  len(arr1) > 1{
		lenAfterComX = len(arr1[1])
	}
	if  len(arr2) > 1{
		lenAfterComY = len(arr2[1])
	}
	//convert to string and remove dot from strings
	strCleanX := strings.Replace(fmt.Sprintf("%g", x), ".", "", -1)
	strCleanY := strings.Replace(fmt.Sprintf("%g", y), ".", "", -1)
	//convert to int
	var _x, _y int
	if intX, err := strconv.Atoi(strCleanX); err == nil {
		_x = intX
	}
	if intY, err := strconv.Atoi(strCleanY); err == nil {
		_y = intY
	}
	//calculate
	var result int
	for i := 0; i < _x; i++  {
		result += _y
	}
	//set dot to string
	var strResult string
	arrResult := strings.Split(strconv.Itoa(result), "")
	for i, v := range arrResult {
		if i == len(arrResult) - (lenAfterComY + lenAfterComX) {
			strResult = strResult + "."
		}
		strResult = strResult + v
	}
	//convert to float64
	answer, err := strconv.ParseFloat(strResult, 64)
	if err != nil {
		fmt.Println(err)
	}
	return answer
}

func fibonacciLoop(n int) int {
	var result int = 0
	value := 1
	if n == 0 {
		return result
	}
	if n == 1 {
		return value
	}
	for i := 0; i < n; i++ {
		result = result + value
		value = result - value 
	}
	return result
}

func fibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == -1 {
		return 1
	}
	if n > 0 {
		return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
	}
	return fibonacciRecursion(n+2) - fibonacciRecursion(n+1)
}

func bubbleSort(arr *[]float32) {
	for i, _ := range *arr {
		for j := 0; j < len(*arr) - i - 1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func uniqueCount(arr *[]int) int {
	count := []int{(*arr)[0]}
	var isTagged bool = false
	for _, v := range *arr {
		for _, newV := range count {
			if v == newV {
				isTagged = true
				break
			}
		}
		if !isTagged {
			count = append(count, v)
		}
		isTagged = false
	}
	fmt.Println(*arr)
	fmt.Println(count)

	return len(count)
}
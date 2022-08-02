package main

import "fmt"

func ShiftArray(array []int, i int) []int {
	temp := []int{}
	if i == 1 {
		index := i + 2
		slice := []int{}
		slice2 := []int{}
		slice3 := []int{}
		slice = append(slice, array[index:index+1]...)
		slice = append(slice, array[:index-1]...)
		slice2 = append(slice2, array[index+3:index+4]...)
		slice2 = append(slice2, array[index+1:index+2]...)
		slice2 = append(slice2, array[index-1:index]...)
		slice2 = append(slice2, array[index:index]...)
		slice3 = append(slice3, array[index+4:]...)
		slice3 = append(slice3, array[index+2:index+3]...)
		temp = append(slice, slice2...)
		temp = append(temp, slice3...)
	}

	if i == 2 {
		index := i + 4
		slice1 := []int{}
		slice2 := []int{}
		slice3 := []int{}
		slice1 = append(slice1, array[index:index+1]...)
		slice1 = append(slice1, array[index-3:index-2]...)
		slice1 = append(slice1, array[index-6:index-5]...)
		slice2 = append(slice2, array[index+1:index+2]...)
		slice2 = append(slice2, array[index-2:index-1]...)
		slice2 = append(slice2, array[index-5:index-4]...)
		slice3 = append(slice3, array[index+2:]...)
		slice3 = append(slice3, array[index-1:index]...)
		slice3 = append(slice3, array[index-4:index-3]...)
		temp = append(slice1, slice2...)
		temp = append(temp, slice3...)
	}
	return temp
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ShiftArray(arr, 2))
}

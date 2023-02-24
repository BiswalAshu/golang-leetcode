package main

import (
	"fmt"
)

func main() {
	// var x int = 10
	// var y int = 20
	// fmt.Println("Previously X = " + strconv.Itoa(x))
	// fmt.Println("Previously Y = " + strconv.Itoa(y))
	// x, y = y, x
	// fmt.Println("Now X = " + strconv.Itoa(x))
	// fmt.Println("Now Y = " + strconv.Itoa(y))
	sliceFromMakeFunc()
}

func MultiValueSingleLineDeclaration() {
	var a, b, c, d = 5, 6.2, 'a', "Ashutosh"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func multiValueSingleBlockDeclartion() {
	var (
		a string = "Ashutosh"
		b int    = 20
	)
	// fmt.Println(a + " " + strconv.Itoa(b))
	fmt.Print(a, " ", b) // print function can take multiple entries comma separated
}

func playArray() { // array are not mutable by length -- fixed size
	var arr = [5]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2, " ", arr[0])
}

func playArray2() {
	arr := [7]int{1: 10} // initialize 1st index as 20 only, leave all other as default
	fmt.Print(arr)
}

func mySlices() { // slice can be mutable by length
	mySlice := []int{}
	mySlice2 := []string{"ja", "to", "gharuku"}
	fmt.Println(mySlice, mySlice2)
	fmt.Println(len(mySlice)) // length of the slice -- how many elements there
	fmt.Println(cap(mySlice)) // capacity of the slice -- how many can stay at max
}

func sliceFromArray() {
	var myArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	mySlice := myArr[2:5]
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}

func sliceFromMakeFunc() []int {
	//using make function to make an slice
	// PRAMETERS --> of []int type, length, capacity as the
	mySlice := make([]int, 3, 8)
	fmt.Println(mySlice)
	return mySlice
}

// []int is a slice
// [10]int in a array
// go loads all elements of a slice to the memory, so make slices with only operational data of an array

/*

} else { --> no error
}
else { --> error

*/

func loops() {
	for i := 10; i > 0; i++ {
		fmt.Println(i)
	}
}
func loops2() {

}

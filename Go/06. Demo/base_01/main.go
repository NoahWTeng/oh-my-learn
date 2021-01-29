package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriteRead() {
	// internal
	fmt.Println("===== INTERNAL PRINT =====")
	print("Used print!\n")
	println("Used printIn!")

	fmt.Println("===== [fmt] PRINT =====")
	// package fmt
	fmt.Print("Used fmt Print! \n")
	fmt.Println("Used fmt Println!")

	fmt.Println("===== [fmt] LITERALS TEMPLATE =====")
	// template literals
	fmt.Printf("Hi, my name is %s, i'm %s years old!\n", "carne", "10")

	// Sprintf formats according to a format specifier and returns the resulting string
	fmt.Printf("name=%s,age=%d,height=%v\n", "Noah", 20, fmt.Sprintf("%.2f", 176.567))

	fmt.Println("===== [fmt] SCAN =====")
	// fmt.Scan
	var name string
	var age int
	fmt.Println("Write name and age")

	// fmt.Scan(&name)
	count, err := fmt.Scan(&name, &age)
	fmt.Println(count, err)
	fmt.Println(name, age)

	// os
	fmt.Println("===== [os] [bufio] =====")
	fmt.Println("Write something...")
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	fmt.Printf("You wrote: %s\n", string(line))
	fmt.Printf("IsPrefix: %s \n", strconv.FormatBool(isPrefix))
	fmt.Printf("Err: %s\n", err)
}

func DeclarationVariable()  {
	// const or var 'name' type = 'value' const need initial value and will never change
	var varName string = "Ian"
	const cname string = "Noah"
	sortName := "Sort"
	fmt.Printf("var name: %s, const name: %s, :=: %s\n ", varName, cname, sortName)

	// multiple
	var varName1, varName2, varName3 string
	const cname1, cname2, cname3 string = "1", "2", "3"
	sortNumber1, sortNumber2, sortNumber3 := 0,1,3
	fmt.Println(varName1, varName2, varName3)
	fmt.Println(sortNumber1, sortNumber2, sortNumber3)

	// _ for no use value
	_, b := 34, 35
	fmt.Printf("not use value _ : b: %s\n", strconv.Itoa(b))
}

func Arrays() {
	// var 'name' [n]type
	var arr_1 [5] int
	fmt.Println(arr_1)
	var arr_2 =  [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr_2)
	arr_3 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr_3)
	arr_4 := [...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)
	arr_5 := [5] int {0:3, 1:5, 4:6} // array position and value
	fmt.Println(arr_5)

	var arr_6 = [3][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 :=  [3][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 :=  [...][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0:3, 1:5, 4:6}, {10, 12, 11, 1, 9}}
	fmt.Println(arr_8)

	arr_9 := make([]int,10)
	fmt.Println(arr_9)

	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(ar)

	a := ar[2:5]
	fmt.Println(a)
}


func changeArray(a []int){
	a[1] = 10
}

func ModifyArray()  {

	/*
	Why does the array value stick?
	func Foo(a [2]int) {
	    a[0] = 8
	}

	func main() {
	    a := [2]int{1, 2}
	    Foo(a)         // Try to change a[0].
	    fmt.Println(a) // Output: [1 2]
	}

	Answer
	Arrays in Go are values.
	When you pass an array to a function, the array is copied.
	*/


	var arr = []int{1,2,3,4,5}
	fmt.Println(arr)
	changeArray(arr)
	fmt.Println(arr)

}



func main()  {
	//WriteRead()
	//DeclarationVariable()
	//Arrays()
	ModifyArray()
}

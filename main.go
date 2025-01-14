package main

import "fmt"

func main()  {
	 a:= [4] int{1,2,3,4}
	 revers(&a)
	 fmt.Println(a)
}

func revers(arr *[4]int )  {
	fmt.Println((arr)[0])
	fmt.Println((arr)[1])
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
		fmt.Println(index, value)
	}
}
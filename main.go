package main

import (
	"fmt"
)

func main() {

	//Read input
	var n int
	fmt.Println("Enter no. of elements:")
	fmt.Scanln(&n)
	b := make([]int, n)

	//Given 1 7 3 45 3 8
	fmt.Println("Enter no. of list of number:")
	for iter := 0; iter < n; iter++ {
		fmt.Scanf("%d", &b[iter])
	}

	//Find largest k
	fmt.Print("The largest number is : ",findKthLargest(b, n),"\n")
	//K's pos
	key,exists := findKthPosition(b, 3)
	if exists {
		fmt.Println("The K value is at index :",key)
	}
	
}

func findKthLargest(slice []int,length int) int {
	for j := 1; j < length; j++ {
		if( slice[0] < slice[j] ) {
			slice[0] = slice[j]
		}
	}
	return slice[0]
}

func findKthPosition(slice []int,target int)(int,bool){
	
	for key,v  := range slice {
		if v == target {
			return key,true
		}
	}
	return 0,false
	
}


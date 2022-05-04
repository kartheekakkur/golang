package main

import "fmt"

func greatestnum(nums []int) (int,int){

	gnum := nums[0]
	j :=0

	for i :=0; i< len(nums);i++{

		if nums[i]>gnum{

			gnum =nums[i]
			j =i
		}

	}

	return gnum,j

	}


func main(){

	nums :=[]int {1,5,3,2,15}

	fmt.Println(greatestnum(nums))
}
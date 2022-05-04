package main

import "fmt"

func selectionsort(nums []int) []int {


for i :=0; i < len(nums)-1;i++{

	lwindex :=i

	for j :=i+1;j <len(nums);j++{
		if nums[j]<nums[lwindex]{
			lwindex =j
		}
	}

	if i !=lwindex{
		temp := nums[i]
		nums[i] = nums[lwindex]
		nums[lwindex] = temp
	}
}	
	return nums

}


func main(){

	nums :=[]int {4,2,1,5,55,3,7}

	fmt.Println(selectionsort(nums))

}

// [4,2,1,5]



package main

import "fmt"

func main()  {

	nums :=[]int{4,2,7,1,3}

	fmt.Println(insertion_sort(nums))
	
}

func insertion_sort(nums []int) []int  {

 	for i:=1; i <len(nums);i++{
       temp_value := nums[i]
       position := i-1

	   for position >=0 {
		if nums[position] >temp_value{
			nums[position+1]=nums[position]
			position -=1
		}else{
             break
		   }
	   }

	   nums[position+1] =temp_value
	}

return nums
	
}
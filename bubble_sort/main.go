package main

import "fmt"


func bubblesort(nums [] int) []int  {
    swapped :=true
	steps :=0

	for swapped{
		steps++
		swapped =false
        for j :=0; j < len(nums)-1; j++ {
			steps++
			if nums[j] >nums[j+1]{
				nums[j+1],nums[j]=nums[j],nums[j+1]
				swapped =true
			}

		}
		
	}
	fmt.Println(steps)

	return nums
	
}


func main()  {

	nums :=[] int{88,87,86,85,84}

	bubblesort(nums)

	fmt.Println(nums)
	
}

// [65,55,45]
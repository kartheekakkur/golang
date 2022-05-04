package main

import "fmt"

// {2,4,6,10}

func twosum(nums []int, target int) []int{
	hm :=make(map[int]int)

	var numlist []int

	for i,v := range nums{
		complement := target - nums[i]

		if j,ok := hm[complement];ok{
           return append(numlist, j,i)
		}
		hm[v]=i
	}

	return numlist
}

func main(){

	nums :=[] int {2,4,6,6}
	target :=5

 fmt.Println(twosum(nums,target))
}
package twosum

func twosum(nums []int, target int) []int  {

	var indices [] int

	for i :=0; i <= len(nums) -2 ; i++{

		if nums[i] + nums[i+1] == target {
			indices =append(indices, i,i+1)
		}
	}


	return indices
	
}
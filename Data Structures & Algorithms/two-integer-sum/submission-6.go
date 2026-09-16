func twoSum(nums []int, target int) []int {
	//Brutual way
   for i := range nums {
		for j:=i+1;j<len(nums);j++ {
			if target-nums[j] == nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil

}

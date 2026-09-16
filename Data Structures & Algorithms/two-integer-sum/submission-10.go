func twoSum(nums []int, target int) []int {
    //Single hash map
	HM := make(map[int]int, len(nums))
	for i := range nums {	
		if j, ok := HM[target-nums[i]]; ok {
			return []int{j, i}
		}
		HM[nums[i]] = i
	}

	return nil
}

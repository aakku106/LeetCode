func hasDuplicate(nums []int) bool {
    for i,vi := range nums{
        for j,vj:= range nums{
            if i!=j{
                if vi==vj{
                    return true
                }
            }
        }
    }
    return false
}

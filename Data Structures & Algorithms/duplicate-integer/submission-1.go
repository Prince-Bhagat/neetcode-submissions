func hasDuplicate(nums []int) bool {
    table := make(map[int]bool)
    for _, ele := range nums{
        _, isExist :=  table[ele]
        if isExist {
            return true
        }else{
            table[ele] = true
        }
    }
    return false
}

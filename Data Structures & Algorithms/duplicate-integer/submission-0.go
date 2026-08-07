func hasDuplicate(nums []int) bool {
    isDuplicate := false
    m := make(map[int]bool)
    for _, element := range nums{
        
        if m[element] {
            isDuplicate = true
            break
        }
        m[element] = true
    }
    return isDuplicate
    
}

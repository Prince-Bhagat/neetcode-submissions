func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for index, value := range nums{
        neededValue := target - value
        itemIndex, isExist := hashMap[neededValue]
        if isExist{
            return []int{itemIndex, index}
        }else{
            hashMap[value]= index
        }
    }
    return []int{}
}

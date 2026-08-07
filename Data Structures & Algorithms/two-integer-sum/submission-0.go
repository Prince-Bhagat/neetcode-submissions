func twoSum(nums []int, target int) []int {
    for i:= 0; i< len(nums)-1;i++{
        numberOne := nums[i]
        for j := i+1;j < len(nums);j++{
            if nums[j] == (target - numberOne){
                return []int{i, j}
            }

        }
    }
    return []int{}
}

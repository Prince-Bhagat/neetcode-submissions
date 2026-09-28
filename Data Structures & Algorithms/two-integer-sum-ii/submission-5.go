func twoSum(numbers []int, target int) []int {

	start := 0
	end := start + 1
	size := len(numbers)
	for end < size{
		sum := numbers[start] + numbers[end]
		for sum <= target {
			if sum == target{
				return []int{start +1, end+1}
			}
			end++
			if end < size {
				sum = numbers[start] + numbers[end]
			}else{
				break
			}
			
		}
		end--

		for (numbers[start] + numbers[end]) <= target{
			if (numbers[start] + numbers[end]) == target {
				return []int{start +1, end +1}
			}
			start++
		}


	}
	return []int{0,0}
}

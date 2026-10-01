func sum(a int, b int, arr []int)int{
	return (arr[a]+ arr[b])
}

func twoSum(input []int, target int)[]int{
	size := len(input)
	start := 0
	end := start +1

	for start < end{
		if end < size && sum(start, end, input) <= target {
			if sum(start, end, input) == target {
				return []int{start +1, end+1}
			}else{
				end++
				continue
			}
			
		}

		end--
		if sum(start, end, input) <= target  {
			
			for sum(start, end, input) <= target{
				if sum(start, end, input) == target {
					return []int{start +1, end +1}
				}else{
					start++
					continue
				}
			}
		}
	}
	return []int{0,0}
} 
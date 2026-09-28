
func longestConsecutive(nums []int) int {
	table := make(map[int]int)
	for _, num := range nums {
		table[num] = num
	}
	// fmt.Println("\ntable ", table)
	
	begining := []int{}

	for _, value := range nums{
		_, isExist := table[value-1]
		if !isExist{
			begining = append(begining, value)
		}
	}	

	// fmt.Println("beginingn ", begining)
	sequences := make(map[int][]int)
	for _, num := range begining {
		// sequences[num] = append(sequences[num], num)
		temp, isExist := table[num]
		if _, isSequenceExist := sequences[num]; isSequenceExist{
			continue
		}
		for isExist {
			// fmt.Printf("\nBefore : Adding %d and isExist is %t. with num %d", temp, isExist, num)
			sequences[num]= append(sequences[num], temp)
			temp = temp +1
			temp , isExist = table[temp]
			// fmt.Printf("\nAfter : Adding %d and isExist is %t", temp, isExist)

		}
	}
	// fmt.Println("\nsequences", sequences)

	length := 0

	for _, val := range sequences{
		if len(val) > length {
			length = len(val)
		}
	}
	return length 

}

type Item struct{
	Map [26]int
	
}

func NewItem(str string)(Item){
	var Map [26]int

	for _, ch := range str{
		Map[ch - 'a']++
		
	}
	return Item{
		Map: Map,
	}
}




func groupAnagrams(strs []string) [][]string {
	array := strs
	table := make(map[[26]int][]string)
	for i := 0; i < len(array); i++ {
		str := array[i]
		
		strItem := NewItem(str)
		value , isExits:= table[strItem.Map]
		if isExits {
			value = append(value, str)
			table[strItem.Map] = value
			
		}else{
			table[strItem.Map] = []string{str}
		}
	}
	var result [][]string
	for _, v := range table{
		result = append(result, v)
	}
	return result
}

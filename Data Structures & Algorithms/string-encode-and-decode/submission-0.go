type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	final := ""
	for _, str := range strs{
		final = final + strconv.Itoa(len(str)) + "#" + str
	}
	return final
}

func (s *Solution) Decode(encoded string) []string {
	final := []string{}
	for counter:= 0;counter < len(encoded);counter++ {
		lengthStr := ""
		ch := encoded[counter]
		for ch != '#' {
			lengthStr = lengthStr + string(ch)
			counter++
			ch = encoded[counter]
		}

		stringStr := ""

		length, _ := strconv.Atoi(lengthStr)
		for i := 0 ;i< length ;i++{
			counter++
			if counter >= len(encoded){
				break
			}
			ch = encoded[counter]
			stringStr = stringStr + string(ch)
			
			
		}
		final = append(final, stringStr)
	}
	return final
}
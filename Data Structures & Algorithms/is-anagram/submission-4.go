func isAnagram(s string, t string) bool {
    frequency := [26]int{}
    for _, sChar := range s{
        frequency[sChar-'a']++
    }
    for _, tChar := range t{
        frequency[tChar -'a']--
    }

    for _, freq := range frequency {
        if freq != 0{
            return false
        }
    }
    return true
}

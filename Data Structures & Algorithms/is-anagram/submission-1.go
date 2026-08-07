func isAnagram(s string, t string) bool {
    hashMap := make(map[rune]int)
    for _, char := range s{
        _, isExist := hashMap[char]
        if(isExist){
            hashMap[char]++
        }else{
            hashMap[char] = 1
        }
    }
    hashMap2 := make(map[rune]int)
    for _, char := range t{
        _, isExist := hashMap2[char]
        if(isExist){
            hashMap2[char]++
        }else{
            hashMap2[char] = 1
        }
    }

    for key, value := range hashMap{
        if value != hashMap2[key]{
            return false
        }
    }
    if (len(hashMap) != len(hashMap2)){
        return false
    }
    return true

}

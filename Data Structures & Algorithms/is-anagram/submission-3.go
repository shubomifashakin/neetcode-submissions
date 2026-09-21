func isAnagram(s string, t string) bool {
    // if the lengths are not the same, return immediately
    if (len(s)!= len(t)){
		return false
	};

    // split both of them 
    runesA := []rune(s)
    runesB:= []rune(t)

    // Sort the rune slice
    sort.Slice(runesA, func(i, j int) bool {
        return runesA[i] < runesA[j]
    })
    sort.Slice(runesB, func(i, j int) bool {
        return runesB[i] < runesB[j]
    })

    // Convert back to string
    sortedStrA := string(runesA)
    sortedStrB := string(runesB)

    answer:= sortedStrA==sortedStrB
    return answer
}
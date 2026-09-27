func isAnagram(s string, t string) bool {
  // we are given 2 strings
  // we need to check if those 2 strings are valie anagrams
  // for them to be valid anagrams they need to be of the same length
  // and contain the same characters
  
  // they need to be of the same length
  if (len(s)!= len(t)){
    return false
  }

    count:=[26]int{}

    // for each character seen, increase the value at its position
    for i:=0;i<len(s);i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    // then if there is a single non zero in the list then we know not all characters are present
    for _,val:=range count {
        // check if the value is a non zero and if it is return false
        if val !=0 {
            return false
        }
    }

    return true
}
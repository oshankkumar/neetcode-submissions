func countSubstrings(s string) int {
    var count int

	expand := func(left,right int){
		for left >=0 && right < len(s) && s[left] == s[right]{
			left--
			right++
			count++
		}
	}

	for i :=0; i < len(s);i++ {
		expand(i,i)
		expand(i,i+1)
	}

	return count
}

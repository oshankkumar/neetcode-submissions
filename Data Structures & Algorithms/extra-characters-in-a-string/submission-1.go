func minExtraChar(s string, dictionary []string) int {
	wordDict := make(map[string]struct{})
	for _, word := range dictionary{
		wordDict[word] = struct{}{}
	}

	memo := make([]int,len(s))
	for i := range memo {
		memo[i] = -1
	}

	return minExtraCharAt(s,memo,0,wordDict)
}

func minExtraCharAt(s string, memo []int,i int, dict map[string]struct{}) int {
	if i == len(s) {
		return 0
	}

	if val := memo[i]; val > 0 {
		return val
	}

	minChar := 1 + minExtraCharAt(s,memo,i+1,dict)

	for w := range dict {
		if strings.HasPrefix(s[i:],w){
			minChar = min(minChar,minExtraCharAt(s,memo,i+len(w),dict))
		}
	}

	memo[i] = minChar
	return minChar
}

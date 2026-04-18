func leastInterval(tasks []byte, n int) int {
	var maxFreq int
	taskFreq := make(map[byte]int)
	for _, b := range tasks {
		taskFreq[b]++
		maxFreq = max(maxFreq, taskFreq[b])
	}

	var maxFreqCount int
	for _, freq := range taskFreq {
		if freq == maxFreq {
			maxFreqCount++
		}
	}

	return max(len(tasks), (n+1)*(maxFreq-1)+maxFreqCount)
}
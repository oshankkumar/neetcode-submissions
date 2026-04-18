from typing import List


class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        freq = dict()
        for n in nums:
            freq[n] = freq.get(n, 0) + 1
            if freq[n] > 1:
                return True
        return False

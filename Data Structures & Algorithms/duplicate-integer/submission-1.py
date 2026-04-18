from typing import List


class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        visted = set()
        for n in nums:
            if n in visted:
                return True
            visted.add(n)
        return False

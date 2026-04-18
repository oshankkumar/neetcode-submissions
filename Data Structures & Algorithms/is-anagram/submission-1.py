from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        c1, c2 = Counter(s), Counter(t)

        return c1 == c2

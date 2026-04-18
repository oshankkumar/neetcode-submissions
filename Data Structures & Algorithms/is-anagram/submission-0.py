from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        c1, c2 = Counter(s), Counter(t)
        if len(c1) != len(c2):
            return False

        for k, c in c1.items():
            if c2.get(k) != c:
                return False

        return True

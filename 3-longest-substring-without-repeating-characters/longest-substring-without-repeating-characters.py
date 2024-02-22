class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        view = set()
        left = 0
        res = 0

        for right in range(len(s)):
            while s[right] in view:
                view.remove(s[left])
                left +=1

            view.add(s[right])
            res = max(res, right-left+1)
        return res
        
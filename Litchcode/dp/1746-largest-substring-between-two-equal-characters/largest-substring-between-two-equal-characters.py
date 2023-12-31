class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        li = []###
        for i in range(len(s)):
            for j in range(i+1,len(s)):
                if s[i]==s[j]:
                    li.append(j-i-1)
        if len(li)==0:
            return -1
        else:
            return max(li)

        
class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        i = 0
        g = sorted(g)
        s = sorted(s)

        for sz in s:
            if i <len(g) and g[i] <=sz:
                i +=1
        return i

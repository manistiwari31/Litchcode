class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        cmin, cmax = 0, 0
        for i, j in zip(s, locked):
            cmin -= 1 if cmin and (j == '0' or i == ')') else -1
            cmax += 1 if j == '0' or i == '(' else -1
            if cmax < 0: return False
        return cmin == 0
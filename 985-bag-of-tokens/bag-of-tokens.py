class Solution:
    def bagOfTokensScore(self, tokens: List[int], power: int) -> int:
        tokens.sort()
        score = 0
        a, b = 0, len(tokens)-1

        while a <= b:
            if power >= tokens[a]:
                power -= tokens[a]
                a+=1
                score+=1
            elif b >= a+1 and score >0:
                power += tokens[b]
                score -= 1
                b -=1
            else:
                break
        return score
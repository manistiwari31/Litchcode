class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        count = {}
        for i in nums:
            if i in count:
                count[i]+=1
            else:
                count[i]=1
        llen = max(count.values())
        ans = []
        for l in range(llen):
            li = []
            for key, value in count.items():
                if value>l:
                    li.append(key)
            if len(li)>0:
                ans.append(li)
        return ans
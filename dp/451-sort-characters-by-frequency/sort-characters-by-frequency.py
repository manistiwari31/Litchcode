class Solution:
    def frequencySort(self, s: str) -> str:
        count = {}
   
        for n in s:
            count[n] = 1 + count.get(n, 0)

        print(count) 
        a = sorted(count.items(), key=lambda x: x[1], reverse=True)    
        print(a)
        st = ""
        for i in a:
            print(i[0])
            st+=i[0]*i[1]

        print(st)

        return st
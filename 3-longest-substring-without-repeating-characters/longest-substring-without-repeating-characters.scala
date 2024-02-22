import scala.collection.mutable
object Solution {
    def lengthOfLongestSubstring(s: String): Int = {
        val view = mutable.HashSet[Char]()
        var left = 0
        var res = 0

        for (right <- s.indices) {
            while(view.contains(s(right))){
                view.remove(s(left))
                left+=1
            } 
            view.add(s(right))
            res = math.max(res,right-left+1)  
        }
        res        
    }
}
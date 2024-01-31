object Solution {
    def containsDuplicate(nums: Array[Int]): Boolean = {
        var hashset = Set[Int]()

        for (n<-nums){
            if (hashset.contains(n)){
                return true
            }

            hashset += n
        }
        false
    }
}
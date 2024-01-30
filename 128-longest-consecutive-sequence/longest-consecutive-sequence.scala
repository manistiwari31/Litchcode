object Solution {
    import scala.collection.mutable
    def longestConsecutive(nums: Array[Int]): Int = {
      val set: mutable.Set[Int] = collection.mutable.Set[Int]()
      nums.foreach(set.add)
      var longestStreak = 0
      for (num<- set)
        if (!set.contains(num - 1)) {
          var currentNum = num
          var currentStreak = 1
          while (set.contains(currentNum + 1)) {
            currentNum += 1
            currentStreak += 1
          }
          longestStreak = math.max(longestStreak, currentStreak)
        }
      longestStreak
    }
}
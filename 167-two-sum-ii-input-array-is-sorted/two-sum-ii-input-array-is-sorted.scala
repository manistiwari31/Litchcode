object Solution {
    def twoSum(numbers: Array[Int], target: Int): Array[Int] = {
    var l = 0
    var r = numbers.length - 1
    while (l < r) {
            var curSum = numbers(l) + numbers(r)

            if (curSum > target) {
                r-=1
            }
            if (curSum < target) {
                l+=1
            }
            if (curSum == target){
                return  Array(l + 1, r + 1)
            }
        }
         Array.empty[Int] 
    }
}
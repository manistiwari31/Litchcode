object Solution {
    def maxArea(height: Array[Int]): Int = {
var left = 0
    var right = height.length - 1
    var max = 0
    while (left < right) {
      val w = right - left
      val h = math.min(height(left), height(right))
      val area = h * w
      max = math.max(max, area)
      if (height(left) < height(right)) left += 1
      else if (height(left) > height(right)) right -= 1
      else {
        left += 1
        right -= 1
      }
    }
    max
  }
}
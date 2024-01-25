object Solution {
    def longestCommonSubsequence(text1: String, text2: String): Int = {
        if (text1.equals(text2)) {
            return text1.length
        }

        val arr1 = text1.toCharArray
        val arr2 = text2.toCharArray

        val dp = Array.ofDim[Int](arr1.length+1, arr2.length+1)

        for (i <- dp.length - 2 to 0 by -1) {
            for (j <- dp.head.length -2 to 0 by -1) {
                if (arr1(i) == arr2(j)) {
                    dp(i)(j) = dp(i+1)(j+1) +1
                } else {
                    dp(i)(j) = Math.max(dp(i)(j+1),dp(i+1)(j))
                }
            }
        }
        dp.head.head
    }
}
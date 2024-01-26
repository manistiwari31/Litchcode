object Solution {
  def findPaths(m: Int, n: Int, maxMove: Int, startRow: Int, startColumn: Int): Int = {
    val mem = Array.fill(m, n, maxMove + 1)(-1L)
    val M = 1000_000_007

    def helper(remainingMoves: Int = maxMove, x: Int = startRow, y: Int = startColumn): Long =
      if (x < 0 || x == m || y < 0 || y == n) 1
      else if (remainingMoves == 0) 0
      else if (mem(x)(y)(remainingMoves) > -1) mem(x)(y)(remainingMoves)
      else {
        mem(x)(y)(remainingMoves) = (helper(remainingMoves - 1, x - 1, y) % M
          + helper(remainingMoves - 1, x + 1, y) % M
          + helper(remainingMoves - 1, x, y - 1) % M
          + helper(remainingMoves - 1, x, y + 1) % M) % M
        mem(x)(y)(remainingMoves)
      }

    (helper() % M).toInt
  }
}

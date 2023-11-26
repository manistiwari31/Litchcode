func largestSubmatrix(matrix [][]int) int {
	last := make([]int, len(matrix[0]))
	maxMatrix := len(matrix[0])
	for col := 0; col < len(matrix[0]); col++ {
		if matrix[0][col] == 0 {
			last[col] = -1
			maxMatrix--
		}
	}
	for row := 1; row < len(matrix); row++ {
		magic := []int{}
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				last[col] = -1
			} else {
				if last[col] == -1 {
					last[col] = row
				}
				magic = append(magic, last[col])
			}
		}
		sort.Ints(magic)
		for i, v := range magic {
			if curMatrix := (i+1)*(row-v+1); curMatrix > maxMatrix {
				maxMatrix = curMatrix
			}
		}
	}
	return maxMatrix
}
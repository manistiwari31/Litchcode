/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func amountOfTime(root *TreeNode, start int) int {
	ans := -1
	graph := getGraph(root)
	q := []int{start}
	seen := make(map[int]bool)
	seen[start] = true

	for len(q) > 0 {
		ans++
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			u := q[0]
			q = q[1:]
			if _, exists := graph[u]; !exists {
				continue
			}
			for _, v := range graph[u] {
				if seen[v] {
					continue
				}
				q = append(q, v)
				seen[v] = true
			}
		}
	}

	return ans
}

func getGraph(root *TreeNode) map[int][]int {
	graph := make(map[int][]int)
	q := [][]*TreeNode{{root, nil}} // {{node, parent}}

	for len(q) > 0 {
		node := q[0][0]
		parent := q[0][1]
		q = q[1:]

		if parent != nil {
			graph[parent.Val] = append(graph[parent.Val], node.Val)
			graph[node.Val] = append(graph[node.Val], parent.Val)
		}

		if node.Left != nil {
			q = append(q, []*TreeNode{node.Left, node})
		}
		if node.Right != nil {
			q = append(q, []*TreeNode{node.Right, node})
		}
	}

	return graph
}
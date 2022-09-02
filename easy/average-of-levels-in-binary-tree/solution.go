/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
  if root == nil { 
    return []float64 {} 
  }
  
  queue := []*TreeNode {root}
  result := []float64{}
  for len(queue) != 0 {
    queueLen, sum := len(queue), 0
    for i := 0; i < queueLen; i++ {
      pop := queue[0]
      queue = queue[1:]
      if pop.Left != nil { 
        queue = append(queue, pop.Left) 
      }
      if pop.Right != nil { 
        queue = append(queue, pop.Right) 
      }
      sum += pop.Val
    }
    result = append(result, float64(sum) / float64(queueLen))
  }
  return result
}

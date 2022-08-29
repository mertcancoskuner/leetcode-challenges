func maxSumSubmatrix(matrix [][]int, k int) int {
  sum := make([][]int, len(matrix)+1)

  for i := range sum {
      sum[i] = make([]int, len(matrix[0])+1)
  }

  for i := range matrix {
      for j := range matrix[i] {
          sum[i+1][j+1] = matrix[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
      }
  }

  result := math.MinInt32
  for i := 1; i <= len(matrix); i++ {
      for j := 1; j <= len(matrix[0]); j++ {
          for p := i; p <= len(matrix); p++ {
              for q := j; q <= len(matrix[0]); q++ {
                  curr := sum[p][q] - sum[i-1][q] - sum[p][j-1] + sum[i-1][j-1]
                  if curr <= k {
                      result = max(result, curr)
                  } 
              }
          }
      }
  }
  return result
}

func max(i, j int) int {
  if i > j {
      return i
  }

  return j
}

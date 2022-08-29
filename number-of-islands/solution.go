func numIslands(grid [][]byte) int {
  var rows = len(grid)
  var columns = len(grid[0])
  var islands = 0
  for row := 0; row < rows; row++ {
    for column := 0; column < columns; column++ {
          if grid[row][column] == '1' {
              islands++
              traverseIsland(grid, row, column)
          }
      }
  }
  return islands
}

func traverseIsland(grid [][]byte, row int, column int ) {
  if(row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) || grid[row][column] == '0') {
      return
  }
  grid[row][column] = '0'
  traverseIsland(grid, row, column+1)
  traverseIsland(grid, row, column-1)
  traverseIsland(grid, row+1, column)
  traverseIsland(grid, row-1, column)
}

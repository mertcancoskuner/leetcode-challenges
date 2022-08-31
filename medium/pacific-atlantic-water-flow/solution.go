func pacificAtlantic(heights [][]int) [][]int {
    pacific := make(map[[2]int]bool)
    for i := 0; i < len(heights[0]); i++ {
        dfs(len(heights)-1, i, -1, &pacific, &heights)
    }
    for i := 0; i < len(heights); i++ {
        dfs(i, len(heights[0])-1, -1, &pacific, &heights)
    }
    
    atlantic := make(map[[2]int]bool)
    for i := 0; i < len(heights[0]); i++ {
        dfs(0, i, -1, &atlantic, &heights)
    }
    for i := 0; i < len(heights); i++ {
        dfs(i, 0, -1, &atlantic, &heights)
    }
    
  intersection := [][]int{}
    for key, _ := range(pacific) {
        if _, ok := atlantic[key]; ok {
            intersection = append(intersection,[]int{key[0], key[1]})
        }
    }
    return intersection
}
        
    
func dfs(row, col, prev_height int, ocean *map[[2]int]bool, heights *[][]int) {
    if row < 0 || row >=len(*heights) {
        return
    }
    if col < 0 || col >= len((*heights)[0]) {
        return
    }
    if _, ok := (*ocean)[[2]int{row, col}]; ok {
        return
    }
    curr_height := (*heights)[row][col]
    if curr_height < prev_height {
        return
    }
    (*ocean)[[2]int{row, col}] = true
    dfs(row-1, col, curr_height, ocean, heights)
    dfs(row+1, col, curr_height, ocean, heights)
    dfs(row, col-1, curr_height, ocean, heights)
    dfs(row, col+1, curr_height, ocean, heights)
}

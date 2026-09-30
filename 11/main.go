package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 发现新岛屿
			if grid[i][j] == '1' {
				count++
				// 用 DFS 把这个岛屿的所有相连陆地都沉没（变成 '0'）
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	// 边界条件检查，如果越界或者当前不是陆地 '1'，则停止搜索
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	// 标记为已访问过，避免死循环
	grid[i][j] = '0'

	// 向上下左右四个方向继续探索
	dfs(grid, i-1, j) // 上
	dfs(grid, i+1, j) // 下
	dfs(grid, i, j-1) // 左
	dfs(grid, i, j+1) // 右
}

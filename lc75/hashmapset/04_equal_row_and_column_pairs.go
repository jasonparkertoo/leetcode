package hashmapset

import "fmt"

func equalPairs(grid [][]int) int {
    pairs := 0
    n := len(grid)

    type row struct {
        vals []int
    }

    colMap := make(map[int]map[string]int)

    // Build column map
    for col := 0; col < n; col++ {
        colVals := make([]int, n)
        for row := 0; row < n; row++ {
            colVals[row] = grid[row][col] 
        }

        if _, ok := colMap[colVals[0]]; !ok {
            colMap[colVals[0]] = make(map[string]int)
        }
        key := fmt.Sprint(colVals)
        colMap[colVals[0]][key]++
    }

    // Match rows against columns
    for row := 0; row < n; row++ {
        rowVals := grid[row]
        if matches, ok := colMap[rowVals[0]]; ok {
            key := fmt.Sprint(rowVals)
            pairs += matches[key]
        }
    }

    return pairs
}
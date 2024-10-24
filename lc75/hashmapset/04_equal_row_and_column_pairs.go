package hashmapset

import "fmt"

// equalPairs takes a n x n grid of integers and returns the number of pairs of equal rows and columns
// The function works by:
// 1. Building a map of column values keyed by their first element for efficient lookup 
// 2. Iterating through each row and checking if it matches any columns in the map
//
// Parameters:
//   grid: A n x n 2D slice of integers
//
// Returns:
//   The number of (row, column) pairs where the elements are equal when reading each top to bottom and left to right
func equalPairs(grid [][]int) int {
    pairs := 0
    n := len(grid)

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
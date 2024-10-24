package hashmapset

import "fmt"

// equalPairs takes a n x n grid of integers and returns the number of pairs of equal rows and columns
// The function works by:
// 1. Building a map of column values keyed by their first element for efficient lookup
// 2. Iterating through each row and checking if it matches any columns in the map
//
// Parameters:
//
//	grid: A n x n 2D slice of integers
//
// Returns:
//
//	The number of (row, column) pairs where the elements are equal when reading each top to bottom and left to right
func equalPairs(grid [][]int) int {
	// Counter for equal row-column pairs
	pairs := 0
	// Size of the grid
	n := len(grid)

	// Map structure: first element of column -> column string representation -> count
	colMap := make(map[int]map[string]int)

	// Build column map
	for col := 0; col < n; col++ {
		// Create slice to store current column values
		colVals := make([]int, n)
		
		// Extract column values from grid
		for row := 0; row < n; row++ {
			colVals[row] = grid[row][col]
		}

		// Initialize inner map if first element not seen before
		if _, ok := colMap[colVals[0]]; !ok {
			colMap[colVals[0]] = make(map[string]int)
		}
		
		// Convert column values to string for map key
		key := fmt.Sprint(colVals)
		
		// Increment count for this column pattern
		colMap[colVals[0]][key]++
	}

	// Match rows against columns
	for row := 0; row < n; row++ {
		// Get current row values
		rowVals := grid[row]
		
		// Check if any columns start with same first element as row
		if matches, ok := colMap[rowVals[0]]; ok {
			// Convert row to string and add matches to pair count
			key := fmt.Sprint(rowVals)
			pairs += matches[key]
		}
	}

	return pairs
}

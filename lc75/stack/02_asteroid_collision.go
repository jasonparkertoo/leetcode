package stack

// astroidCollision takes a slice of integers representing asteroids, where positive values
// indicate asteroids moving right and negative values indicate asteroids moving left.
// The absolute value indicates the size of the asteroid.
//
// When two asteroids collide:
//	- If they are the same size, both asteroids are destroyed
//	- If one is larger, only the smaller asteroid is destroyed
//	- Asteroids only collide if they are moving towards each other (right meets left)
//
// Time Complexity: O(n) where n is the length of the asteroid array
//	- In the worst case, we may need to process each asteroid once
//	- While there is a nested loop, each asteroid can only be popped from result at most once
// 
// Space Complexity: O(n) 
// 	- We use a slice to store the result which in the worst case may need to store all asteroids
//	- No other data structures are used
func astroidCollision(asteroids []int) []int {
	result := make([]int, 0)

	// Iterate through each asteroid in the array
	// For each asteroid, check for collisions with previous asteroids
	// currAsteroid can either:
	//	- Get destroyed if it hits a larger asteroid
	//	- Destroy other asteroids if it's larger 
	//	- Pass through if it never collides with another asteroid
	//	- Be destroyed along with another asteroid if they collide and are the same size
	for _, currAsteroid := range asteroids {
		isAsteroidDestroyed := false

		// While we have asteroids in result and current asteroid is going left 
		// and last asteroid in result is going right
		for len(result) > 0 && currAsteroid < 0 && result[len(result)-1] > 0 {
			lastAsteroid := result[len(result)-1]
			currAsteroidSize := -currAsteroid

			// If last asteroid is smaller, it gets destroyed
			if lastAsteroid < currAsteroidSize {
				result = result[:len(result)-1] 
				continue
			}

			// If asteroids are same size, both get destroyed
			if lastAsteroid == currAsteroidSize {
				result = result[:len(result)-1]
				isAsteroidDestroyed = true
			} else {
				// Last asteroid is bigger, current asteroid gets destroyed
				isAsteroidDestroyed = true
			}
			break
		}

		if !isAsteroidDestroyed {
			result = append(result, currAsteroid)
		}
	}

	return result
}

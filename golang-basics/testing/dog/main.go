// Package dog allows to understand more dogs.
package dog

// Years confverts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog Years
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

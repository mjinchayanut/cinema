package movie

import "fmt"

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating: %.1f\n", name, rating)
}

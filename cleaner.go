package cleaner

import (
	"log"
)

// Cleaner represents a module to be cleaned
type Cleaner struct {
	// Name is the semi-representative name of the
	// cleaner
	Name string

	// CleanUp is the cleaner function
	CleanUp func()
}

// Cleaners is the list of Cleaners
var Cleaners []Cleaner

// AddCleaner appends to the cleaners array
func AddCleaner(c Cleaner) {
	Cleaners = append(Cleaners, c)
}

// CleanUp runs all cleaners iteratively,
// in their separate goroutine
func CleanUp() {
	log.Println("Cleaning up")

	for i := 0; i < len(Cleaners); i++ {
		c := Cleaners[len(Cleaners)-i-1]

		c.CleanUp()
		log.Println(c.Name, "cleaned up")
	}
	log.Println("Cleaned up")
}

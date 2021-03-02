package cleaner

import (
	"log"
	"sync"
)

// Cleaner represents a module to be cleaned
type Cleaner struct {
	// Name is the semi-representative name of the
	// cleaner
	Name string

	// CleanUp is the cleaner function
	CleanUp func()
}

// cleaners is the list of cleaners
var cleaners []Cleaner

// AddCleaner appends to the cleaners array
func AddCleaner(c Cleaner) {
	cleaners = append(cleaners, c)
}

// CleanUp runs all cleaners iteratively,
// in their separate goroutine
func CleanUp() {
	var wg sync.WaitGroup
	wg.Add(len(cleaners))

	log.Println("Cleaning up")

	for i := 0; i < len(cleaners); i++ {
		go func(i int) {
			defer wg.Done()
			c := cleaners[i]

			c.CleanUp()
			log.Println(c.Name, "cleaned up")
		}(i)
	}
	wg.Wait()
	log.Println("Cleaned up")
}

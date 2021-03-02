package cleaner_test

import (
	"testing"

	"github.com/manen/cleaner"
)

func Test(t *testing.T) {
	cleaner.AddCleaner(cleaner.Cleaner{
		Name: "Test",
		CleanUp: func() {
		},
	})
	cleaner.AddCleaner(cleaner.Cleaner{
		Name: "Bruh",
		CleanUp: func() {
		},
	})

	cleaner.CleanUp()
}

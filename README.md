# Cleaner

Cleaner for Go

## Installation

```sh
go get github.com/manen/cleaner
```

## Usage

```go
package main

import (
  "github.com/manen/cleaner"
)

func main() {
  // Create a new cleaner
  c := cleaner.Cleaner{
    Name: "Example",
    CleanUp: func() {
      // Clean up
    }
  }

  // Add the new cleaner
  cleaner.AddCleaner(c)

  // Run all the cleaners
  cleaner.CleanUp()
}
```

## License

[GPL 3](LICENSE.txt)

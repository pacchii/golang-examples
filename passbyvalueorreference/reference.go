package passbyvalueorreference

import "fmt"

func Example() {
	fmt.Println("Go is always pass-by-value. No exceptions. ---BUT--- Some types (slice, map, channel, pointer, function) are reference-like values.")
}

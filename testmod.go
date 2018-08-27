package testmod

import "fmt"

// Hi Says Hi to someone
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}

package dog

import (
	"fmt"
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

func From11() {
	fmt.Println("Dog repo v1.1.0")
}

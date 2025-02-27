package main

import (
	"fmt"
)

func main() {
	//Define shark variable as a slice of strings
	sharks := []string{"hammerhead", "great white", "dog fish", "frilled"}
	// For loop that iterates over sharks list ans prints each item
	for _, shark := range sharks {
		fmt.Println(shark)
	}
}

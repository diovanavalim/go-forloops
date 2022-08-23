package forloops

import "fmt"

func ForRange(slice []string) {
	fmt.Println("For Range")
	for index, word := range slice {
		fmt.Printf("Index %d, palavra %s\n", index, word)
	}
}

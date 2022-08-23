package forloops

import "fmt"

func ForCondition(slice []string) {
	var sliceLen int = len(slice)
	var counter int = 0

	fmt.Println("For Condition")
	for counter < sliceLen {
		fmt.Printf("Elemento %d: %s\n", counter, slice[counter])
		counter++
	}
}

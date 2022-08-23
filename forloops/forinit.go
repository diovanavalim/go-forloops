package forloops

import "fmt"

func ForInit(slice []string) {
	var sliceLen int = len(slice)

	fmt.Println("For Init")
	for i := 0; i < sliceLen; i++ {
		fmt.Printf("Índice %d, elemento %s\n", i, slice[i])
	}
}

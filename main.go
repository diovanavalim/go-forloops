package main

import (
	"forloop/forloops"
)

func main() {
	slice1 := []string{"Batata", "Maçã", "Cenoura", "Beterraba", "Alface", "Pepino"}
	slice2 := []string{"Maria", "Tânia", "Joana", "Beatriz", "Ana", "Bárbara", "Camila"}
	slice3 := []string{"Apple", "Samsung", "Xiaomi", "Motorola", "Dell", "HP"}

	forloops.ForCondition(slice1)
	forloops.ForInit(slice2)
	forloops.ForRange(slice3)
}

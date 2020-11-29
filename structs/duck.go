package structs

import "fmt"

type Duck struct {
	Name string
}

func (d Duck) Breathe() {
	fmt.Println("Duck breathes")
}

func (d Duck) Walk() {
	fmt.Println(d.Name)
	fmt.Println("Duck walk")
}

func (d Duck) Speed() int {
	return 0
}

package structs

import "fmt"

type Lion struct {
	Age int
}

func (l Lion) Breathe() {
	fmt.Println("Lion breathes")
}

func (l Lion) Walk() {
	fmt.Println(l.Age)
	fmt.Println("Lion walk")
}

func (l Lion) Speed() int {
	return 0
}

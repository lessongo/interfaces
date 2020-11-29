package main

import (
	"github.com/lessongo/interface/services"
	"github.com/lessongo/interface/structs"
)

func main() {
	lion := structs.Lion{Age: 10}
	services.CallBreathe(lion)
	services.CallWalk(lion)
	services.CallSpeed(lion)

	duck := structs.Duck{Name: "Cingozr"}
	services.CallBreathe(duck)
	services.CallWalk(duck)
	services.CallSpeed(duck)
}

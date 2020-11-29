package services

import "github.com/lessongo/interface/interfaces"

func CallBreathe(a interfaces.Animal) {
	a.Breathe()
}

func CallWalk(a interfaces.Animal) {
	a.Walk()
}

func CallSpeed(a interfaces.Animal) {
	a.Speed()
}

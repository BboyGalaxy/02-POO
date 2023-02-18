package main

import (
	"fmt"

	"github.com/BboyGalaxy/02-POO/course"
)

func main() {
	Go := course.Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	/*css := Course{
		Name:   "Css desde cero",
		IsFree: true,
	}

	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}*/

	fmt.Println(Go.Name)
	//fmt.Printf("%+v\n", css)
	//fmt.Printf("%+v", js)
	Go.PrintClasses()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)

}

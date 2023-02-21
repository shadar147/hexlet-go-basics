package ex26

import (
	"fmt"
)

func Test() {
	fmt.Println(CopyParent(nil))

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
			{
				Name: "Vasya",
				Age:  22,
			},
		},
	}

	cp := CopyParent(p)
	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{}
	fmt.Println(p.Children) // [{Andy 18}]
	fmt.Println(p, cp)
}

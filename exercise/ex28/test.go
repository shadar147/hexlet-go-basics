package ex28

import "fmt"

func Test() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	fmt.Println(pl.GetAgePopularity()) // map[18:2 44:1]

	newPl := PersonList{}

	fmt.Println(newPl.GetAgePopularity())
}

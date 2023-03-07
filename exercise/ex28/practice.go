package ex28

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int, len(pl))

	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

package ex31

import "fmt"

func Test() {
	job, err := ExecuteMergeDictsJob(&MergeDictsJob{})

	fmt.Printf("%+v\n", job)
	fmt.Println(err)

	job, err = ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
		nil,
	}})

	fmt.Printf("%+v\n", job)
	fmt.Println(err)

	job, err = ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b", "b": "c"},
		{"d": "e", "f": "g"},
		{"a": "z", "f": "g"},
	}})

	fmt.Printf("%+v\n", job)
	fmt.Println(err)
}

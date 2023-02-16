package ex24

import "fmt"

func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%0.2f in my wallet right now.", name, age, money)
}

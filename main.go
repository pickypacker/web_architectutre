package main

// Person data structure
type Person struct {
	First string
}

func main() {

	p := Person{First: "Amie"}

	PrintFirst(p)

}

// PrintFirst print the Person structure First value
func PrintFirst(p Person) string {

	println(p.First)

	return p.First
}

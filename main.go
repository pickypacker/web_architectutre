package main

type Person struct {
	First string
}

func main() {

	p := Person{First: "Amie"}

	PrintFirst(p)

}

func PrintFirst(p Person) {

	println(p.First)
}

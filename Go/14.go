package main

import "fmt"

/* 1
// interface
type Printer interface {
	Print() string
}

func Print(p Printer) {
	fmt.Println(p.Print())
}

// types
type Message struct {
	name string
}

func (m Message) Print() string {
	return fmt.Sprintf("%v from Print method", m.name)
}

func main() {
	var input string
	fmt.Scan(&input)

	mess := Message{
		name: input,
	}

	Print(mess)
}
*/

// interface
type Stringer interface {
	Length() int
}

func Length(s Stringer) int {
	return s.Length()
}

// types
type MyString string

func (m MyString) Length() int {
	return len(m)
}

func main() {
	var input string
	fmt.Scan(&input)

	mess := MyString(input)

	fmt.Println("Length of the string:", Length(mess))
}

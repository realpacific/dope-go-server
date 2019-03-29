package main

import "fmt"

func main() {
	c := Cat{"Tom"}
	fmt.Println(c.speak())

	d := Dog{"Rufus"}
	fmt.Println(d.speak())

	pd := PythonDeveloper{"Python Develoer"}
	fmt.Println(pd.speak())

	separator()

	// Go dynamic arrays aka "slice" 
	animals := []Animal{Dog{"Snoop"}, Cat{"Meowt"}, PythonDeveloper{"SnakeEyes"}}
	for i, animals := range animals {
		fmt.Printf(fmt.Sprintf("%d. %s", i+1, animals.speak()))
		fmt.Printf("\n")
	}

	separator()

}


func separator() {
	count := 0
	for count < 30 {
		fmt.Print("-")
		count++
	}
	fmt.Println("-")
		

}

/**
 * we define Animals as any type that has a method named `speak() string`
**/
type Animal interface {
	speak() string
}

type Cat struct{
	name string
}

func (c Cat) speak() string {
	return fmt.Sprintf("%s says \"Meow\"", c.name)
} 

type Dog struct {
	name string
}

func (d Dog) speak() string {
	return fmt.Sprintf("%s says \"Woof Woof\"", d.name)
}

type PythonDeveloper struct {
	name string
}

func (pd PythonDeveloper) speak() string {
	return fmt.Sprintf("%s says \"Crap! I missed a space\"", pd.name)
}
package main

import (
	pt "github.com/alessandrobessi/piecetable/pkg/piecetable"
	"fmt"
)

func main() {

	// read from file using buffer chunks of 1024 bytes
	p := pt.ReadFromFile("../data/numbers.txt", 1024)
	fmt.Println(p.GetText())
	
	// print only some specific lines
	for _, i := range []int{0, 1, 2, 9} {
		fmt.Println(p.GetLine(i))
	}

	p = pt.ReadFromFile("../data/pangram.txt", 1024)
	fmt.Println(p.GetText())

	// add "Foxy, " at the beginning
	p.Insert("Foxy, ", 0) 
	fmt.Println(p.GetText())

	// replace "quick" with "lazy"
	p.Delete(10 , 6) 
	p.Insert("lazy ", 10)
	fmt.Println(p.GetText())

}

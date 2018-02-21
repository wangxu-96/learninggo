package main

import (
	"fmt"
)

func Grade(score int) {
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("%d\n", score))
	case score < 60:
		fmt.Println("E")
	case score < 70:
		fmt.Println("D")
	case score < 80:
		fmt.Println("C")
	case score < 90:
		fmt.Println("B")
	case score < 100:
		fmt.Println("A")
	}
}

func main() {
	/*const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", contents)
	}*/

	Grade(59)
	Grade(69)
	Grade(79)
	Grade(89)
	Grade(99)
	Grade(109)

}

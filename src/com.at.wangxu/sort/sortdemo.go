package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [] int{3,4,5,6,7,1}
	sort.Ints(a)
	for _,v:= range a {
		fmt.Println(v)
	}
}

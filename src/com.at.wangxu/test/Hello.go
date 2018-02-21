package main

import (
	"fmt"
	"math"
)


var (
 s=123;

)

func initValue(){
	s :="sss"
	fmt.Println(s)

}
func oular(){
	var a,b int=3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

const(
	a=1<<(10*iota)
	b
	c
	d
)

func enums()  {
	fmt.Println(a,b,c,d)
}
func main() {
	fmt.Println("HelloWorld")
	initValue()
	oular()
	enums()
}

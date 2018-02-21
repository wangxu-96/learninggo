package main

import (
	"learninggo/src/pipeline"
	"fmt"
)

func main() {
	p :=
		pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(1, 3, 5, 7, 9, 0)), pipeline.InMemSort(pipeline.ArraySource(2, 4, 6, 8, 10, 1)))
	for v := range p {
		fmt.Println(v)
	}
}

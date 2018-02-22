package main

import (
	"learninggo/src/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename  = "Lagare.in"
	const count  =50
	file,err:=os.Create(filename)
	if err!=nil {
		panic(err)
	}
	defer file.Close()
	p:=pipeline.RandomSource(count)
	buffer:=bufio.NewWriter(file)
	pipeline.WriteSink(buffer,p)
	buffer.Flush()
	file,err=os.Open(filename)
	if err!=nil {
		panic(err)
	}
	defer file.Close()
	p=pipeline.ReaderSource(bufio.NewReader(file))
	for v:=range p{
		fmt.Println(v)
	}
}
func MergeDemo(){
	p :=
		pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(1, 3, 5, 7, 9, 0)), pipeline.InMemSort(pipeline.ArraySource(2, 4, 6, 8, 10, 1)))
	for v := range p {
		fmt.Println(v)
	}
}
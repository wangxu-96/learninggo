package main

import (
	"os"
	"learninggo/src/pipeline"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	p:=createNetworkPipeline("Small.in", 512, 4)
	WriteToFile(p, "Small.out")
	printFile("Small.out")
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}
func WriteToFile(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriteSink(writer, p)
}
func createPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sourceResults := [] <-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sourceResults = append(sourceResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sourceResults...)
}
func createNetworkPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	sourceResults:=[]<-chan int{}
	for _,addr := range sortAddr{
		sourceResults = append(sourceResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sourceResults...)
}

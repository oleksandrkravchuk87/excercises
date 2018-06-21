package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./text")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("line:%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file: %s", err)
	}

	//scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords)
	//
	//// This is our buffer now
	//var lines []string
	//
	//for scanner.Scan() {
	//	lines = append(lines, scanner.Text())
	//}
	//
	//fmt.Println("read lines:")
	//for _, line := range lines {
	//	fmt.Println("+++", line)
	//}
}

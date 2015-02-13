package main

import (
	"fmt"
	// "io/ioutil"
	"strings"
	"os"
	"bufio"
)

func main() {
	// content, err := ioutil.ReadFile("testFiles/fileRead2.txt")
	// if err != nil {
	//     //Do something
	// }
	// lines := strings.Split(string(content), "\n")
	// fmt.Println(lines)
	// fmt.Println(string(content))
	readLine("testFiles/fileRead2.txt")
}

func readLine(path string) {
  inFile, _ := os.Open(path)
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines) 
  
  for scanner.Scan() {
  	text := strings.Fields(scanner.Text())
  	fmt.Println(text)
  }
}

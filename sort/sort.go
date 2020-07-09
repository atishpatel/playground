package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read all at same time
	fileData, err := ioutil.ReadFile("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	array, err := readArray(fileData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(array)
	// read line by line
	arr, err := readLineByLine("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
}

func readLineByLine(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	array := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		for j := range numbers {
			val, err := strconv.ParseInt(numbers[j], 10, 64)
			if err != nil {
				return nil, err
			}
			array = append(array, int(val))
		}
	}
	return array, nil
}

func readArray(fileData []byte) ([]int, error) {
	lines := strings.FieldsFunc(string(fileData), func(r rune) bool {
		return r == '\n' || r == '\r'
	})
	array := []int{}
	for i := range lines {
		numbers := strings.Fields(lines[i])
		for j := range numbers {
			val, err := strconv.ParseInt(numbers[j], 10, 64)
			if err != nil {
				return nil, err
			}
			array = append(array, int(val))
		}
	}

	return array, nil
}

func sortBuiltIn(array []int) []int {
	sort.Ints(array)
	return array
}

// merge sort

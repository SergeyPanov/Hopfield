package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"Hopfield/hopfield"
	"flag"
)

func readContent(path string) [][]int {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		print("File ", path, " does not exists.")
		os.Exit(1)
	}
	var inputs [][]int
	strs := strings.Split(string(bs), "\n")
	for i := 0; i < len(strs); i++ {
		numbers := strings.Split(strs[i], " ")

		var convertedNumbers []int

		for n := 0; n < len(numbers); n++ {
			converted, err := strconv.Atoi(numbers[n])
			if err != nil {
				print("Invalid file")
				print(err.Error())
				os.Exit(1)
			}
			convertedNumbers = append(convertedNumbers, converted)
		}
		inputs = append(inputs, convertedNumbers)
	}
	return inputs

}

func main() {

	ideal := flag.String("ideal", "", "Path to file with ideal vectors.")
	input := flag.String("input", "", "Path to file with input damaged files")
	flag.Parse()

	net := new(hopfield.Network)
	net.Setup(readContent(*ideal))

	damaged := readContent(*input)

	for i := 0; i < len(damaged); i++ {
		net.Restore(damaged[i])
	}

}

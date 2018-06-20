package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"Hopfield/hopfield"
	"fmt"
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
	const input = "/Users/sergeypanov/go/src/Hopfield/resources/input"

	mt1 := new(hopfield.Matrix)
	mt2 := new(hopfield.Matrix)

	mt1.Init(readContent(input))
	mt2.Init(readContent(input))



	mt2.Transpose()

	fmt.Println(mt2.Matrix)
	fmt.Println(mt1.Matrix)

	mt2.Mult(mt1)

	fmt.Println(mt2)





}
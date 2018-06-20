package hopfield

import (
	"fmt"
)

type Network struct {
	Ideal         Matrix
	Weights       Matrix
	Input, Output []int
	Threshold     float64
}

func (net *Network) Setup(images [][]int) {

	// Store ideal images
	net.Ideal = Matrix{}
	net.Ideal.Init(images)

	// Setup initial weights
	net.Weights = Matrix{}
	net.Weights.Init(images)
	net.Weights = *net.Weights.Transpose()
	net.Weights = *net.Weights.MultByMatrix(&net.Ideal)
	for i := 0; i < len(net.Weights.Matrix); i ++ {

		for j := 0; j < len(net.Weights.Matrix[0]); j ++ {
			if i == j {
				net.Weights.Matrix[i][j] = 0
			}
		}
	}

	net.Output = make([]int, len(images[0]))
	net.Input = make([]int, len(images[0]))
}

// Activation function
func (net *Network) activation(states []int) []int {
	for i := 0; i < len(states); i++ {
		if states[i] <= int(net.Threshold) {
			states[i] = -1
		} else {
			states[i] = 1
		}
	}
	return states
}

// Zero distance condition
func (net *Network) distance(prev, states []int) float64 {
	eulerDistance := 0
	for i := 0; i < len(states); i ++ {
		eulerDistance += (states[i] - prev[i]) * (states[i] - prev[i])
	}
	return float64(eulerDistance)

}

func (net *Network) Restore(input []int) {
	// First iteration Input == Output
	net.Output = input

	iteration := 0

	output_2 := make([]int, len(input))
	output_3 := make([]int, len(input))

	for true {
		states := net.Weights.MultByVector(net.Output)
		if iteration%2 == 0 {
			output_3 = net.Output
		} else {
			output_2 = net.Output
		}

		net.Input = net.Output
		net.Output = net.activation(states)

		if net.distance(net.Input, net.Output) == 0 {
			break
		}

		if net.distance(net.Output, output_2) == 0 && net.distance(net.Input, output_3) == 0 {
			fmt.Println("Image was not restored. Stop by altration.")
			break
		}
		iteration += 1

	}

	fmt.Println(net.Output)

}

package hopfield

type Matrix struct {
	Matrix [][]int
}

// Init matrix by input matrix "inM"
func (m *Matrix) Init(inM [][]int)  {
	colS := len(inM)
	for i := 0; i < colS; i++ {
		m.Matrix = append(m.Matrix, inM[i])
	}
}

// Transpose matrix
func (m *Matrix) Transpose()  {
	newMatrix := new(Matrix)
	newMatrix.Matrix = make([][]int, len(m.Matrix[0]))
	rSize := len(m.Matrix)
	cSize := len(m.Matrix[0])

	for i := 0; i < len(m.Matrix[0]); i++ {
		newMatrix.Matrix[i] = make([]int, rSize)
	}

	for i := 0; i < rSize; i++ {

		for j := 0; j < cSize; j++ {
			newMatrix.Matrix[j][i] = m.Matrix[i][j]
		}
	}
	m.Matrix = newMatrix.Matrix
}

// Multiply 2 matrix
func (m *Matrix) Mult(mt *Matrix){

	resMatrix := new(Matrix)

	resMatrix.Matrix = make([][]int, len(m.Matrix))

	for i := 0; i < len(resMatrix.Matrix) ; i ++ {
		resMatrix.Matrix[i] = make([]int, len(mt.Matrix[0]))
	}

	for i := 0; i < len(resMatrix.Matrix) ; i++ {

		for j := 0; j < len(mt.Matrix[0]); j ++{
			sum := 0

			for k := 0; k < len(mt.Matrix); k ++{
				sum += m.Matrix[i][k] * mt.Matrix[k][j]
			}
			resMatrix.Matrix[i][j] = sum
		}
	}


	m.Matrix = resMatrix.Matrix


}
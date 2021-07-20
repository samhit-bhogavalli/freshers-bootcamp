package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	rows int
	columns int
	values [][]int
}

//returns no. of rows
func (m Matrix) getRow() int {
	return m.rows
}

//returns no.of columns
func (m Matrix) getColumns() int {
	return m.columns
}

//sets the value at i, j position
func (m Matrix) setValueAtIJ(i, j, value int) {
	if i >= m.rows || j >= m.columns {
		fmt.Println("out of bounds")
		return
	}
	m.values[i][j] = value
}


//creates the matrix of required size and initialises it
func (m *Matrix) init(row, column int){
	m.rows = row
	m.columns = column
	m.values = make([][]int, row)
	for i := 0; i < column; i++ {
		m.values[i] = make([]int, column)
	}
}

//creates a new matrix of max row and max column and adds the two matrices
func (mat1 Matrix) add(mat2 Matrix) Matrix{

	maxRow, maxColumn := int(math.Max(float64(mat1.rows), float64(mat2.rows))), int(math.Max(float64(mat1.columns), float64(mat2.columns)))
	var result = make([][]int, maxRow)
	for i := 0; i < maxRow; i++ {
		result[i] = make([]int, maxColumn)
	}
	minRow, minColumn := int(math.Min(float64(mat1.rows), float64(mat2.rows))), int(math.Min(float64(mat1.columns), float64(mat2.columns)))
	for i := 0; i < minRow; i++ {
		for j := 0; j < minColumn; j++ {
			result[i][j] = mat1.values[i][j] + mat2.values[i][j]
		}
	}
	mat3 := Matrix{maxRow, maxColumn, result}
	return mat3
}
func main(){
	//testing
	mat := Matrix{0, 0,[][]int{}}
	mat.init(5,5)
	fmt.Println(mat)
	mat.setValueAtIJ(1,1,1)
	mat2 := Matrix{0, 0,[][]int{}}
	mat2.init(2,2)
	fmt.Println(mat2)
	fmt.Println(mat2.add(mat))
}

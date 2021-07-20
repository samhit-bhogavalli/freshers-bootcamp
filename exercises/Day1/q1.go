package main

import (
	"errors"
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
func (m Matrix) setValueAtCoordinates(row, column, value int) error {
	if row >= m.rows || column >= m.columns {
		return errors.New("out of bounds")
	}
	m.values[row][column] = value
	return nil
}


//creates the matrix of required size and initialises it
func (m *Matrix) init(){
	m.values = make([][]int, m.rows)
	for i := 0; i < m.rows; i++ {
		m.values[i] = make([]int, m.columns)
	}
}

//creates a new matrix of max row and max column and adds the two matrices
func (mat1 Matrix) add(mat2 Matrix) Matrix{
	maxRow, maxColumn := int(math.Max(float64(mat1.rows), float64(mat2.rows))), int(math.Max(float64(mat1.columns), float64(mat2.columns)))
	var result = Matrix{maxRow, maxColumn, [][]int{}}
	result.values = make([][]int, maxRow)
	for i := 0; i < maxRow; i++ {
		result.values[i] = make([]int, maxColumn)
	}
	minRow, minColumn := int(math.Min(float64(mat1.rows), float64(mat2.rows))), int(math.Min(float64(mat1.columns), float64(mat2.columns)))
	for i := 0; i < minRow; i++ {
		for j := 0; j < minColumn; j++ {
			result.values[i][j] = mat1.values[i][j] + mat2.values[i][j]
		}
	}
	return result
}

func main(){
	//testing
	var mat = Matrix{3, 3, [][]int{{1,2,3},{4,5,6},{7,8,9}}}
	//mat.init()
	fmt.Println(mat)
	if er := mat.setValueAtCoordinates(5,1,1); er != nil {
		fmt.Println(er)
	}
	fmt.Println(mat)
	mat2 := Matrix{3, 3, [][]int{{1,2,3},{4,5,6},{7,8,9}}}
	//mat2.init()
	fmt.Println(mat2)
	fmt.Println(mat2.add(mat))
}

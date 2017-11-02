package matrix

import (
	"fmt"
	"log"
	"os"
)

type Matrix struct {
	rows    int       // 行
	columns int       // 列
	matrix  []float64 // 行 * 列の長さ
}

func NewMatrix(row, column int) *Matrix {
	matrix := new(Matrix)
	matrix.rows = row
	matrix.columns = column
	matrix.matrix = make([]float64, matrix.rows*matrix.columns)
	return matrix
}

// checkBroken look the length of row and column if t
func (m *Matrix) checkBroken() {
	if m.rows <= 0 || m.columns <= 0 {
		log.Fatal("Matrix is broken")
		os.Exit(1)
	}
}

// 四則演算
func (m *Matrix) Add(num interface{}) {
	m.checkBroken()
	if mat, ok := num.(Matrix); ok {
		if mat.rows != m.rows || mat.columns != m.columns {
			log.Fatal("The row and column num are different")
			os.Exit(1)
		}
		for i, val := range mat.matrix {
			m.matrix[i] += val
		}
		return
	}
	if mat, ok := num.(*Matrix); ok {
		if mat.rows != m.rows || mat.columns != m.columns {
			log.Fatal("The row and column num are different")
			os.Exit(1)
		}
		for i, val := range mat.matrix {
			m.matrix[i] += val
		}
		return
	}
	if mat, ok := num.(int); ok {
		for i := range m.matrix {
			m.matrix[i] += float64(mat)
		}
		return
	}
	if mat, ok := num.(float64); ok {
		for i := range m.matrix {
			m.matrix[i] += mat
		}
		return
	}
	log.Fatal("The add type is not collect")
	os.Exit(1)
}

// 行列の積

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() {
	m.checkBroken()
	m.matrix = make([]float64, m.rows*m.columns)
}

func (m *Matrix) initByArray(row []float64) {
	m.rows = 1
	m.columns = len(row)
	m.matrix = row
}

// AddRow add row if the len of column = 0. create new vector 1 * len(row)
func (m *Matrix) AddRow(row []float64) {
	if m.columns != len(row) && m.columns != 0 {
		log.Fatal("Column length is not same")
		os.Exit(1)
	}
	if m.columns == 0 {
		m.initByArray(row)
		return
	}
	m.rows++
	m.matrix = append(m.matrix, row...)
}

func (m *Matrix) AddRowHEAD(row []float64) {
	if m.columns != len(row) && m.columns != 0 {
		log.Fatal("Column length is not same")
		os.Exit(1)
	}
	if m.columns == 0 {
		m.initByArray(row)
		return
	}
	m.rows++
	m.matrix = append(row, m.matrix...)
}

// Show will show matrix condition
func (m *Matrix) Show() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			fmt.Print(m.matrix[i*m.columns+j])
		}
		fmt.Println()
	}
}

// Size show this matrix size
func (m *Matrix) Size() {
	fmt.Println(m.rows, " ", m.columns)
}

func (m *Matrix) checkSize(row, column int) {
	if row <= 0 || column <= 0 || row > m.rows || column > m.columns {
		log.Fatal("Size is invalid")
		os.Exit(1)
	}
}

// At show a point of value
func (m *Matrix) At(row, column int) {
	m.checkSize(row, column)
	fmt.Println(m.matrix[column*(row-1)+column-1])
}

// Set will set specifix value
func (m *Matrix) Set(row, column int, value float64) {
	m.checkSize(row, column)
	m.matrix[column*(row-1)+column-1] = value
}

// TODO: 転置行列の作成

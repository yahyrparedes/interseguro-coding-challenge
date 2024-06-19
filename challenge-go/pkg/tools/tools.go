package tools

import (
	"fmt"
	"math"
)

func MultiplyMatrices(A, B [][]float64) [][]float64 {
	rowsA := len(A)
	colsA := len(A[0])
	rowsB := len(B)
	colsB := len(B[0])
	if colsA != rowsB {
		panic("El numero de columnas debe ser igual al numero de filas de la matriz B")
	}

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}

func QRFactorization(matrix [][]float64) ([][]float64, [][]float64) {
	// Apply the Gram-Schmidt process to get Q
	Q := ApplyGramSchmidt(matrix)
	fmt.Println("Q", Q)
	// Calculate R as the product of Q transpose and the original matrix
	QTranspose := Transpose(Q)
	fmt.Println("QTranspose", QTranspose)
	R := MultiplyMatrices(QTranspose, matrix)

	return Q, R
}

func InvertMatrixRectangle(data [][]float64) [][]float64 {
	var n [][]float64
	for i := 0; i < len(data[0]); i++ {
		var row []float64
		for j := 0; j < len(data); j++ {
			row = append(row, data[j][i])
		}
		n = append(n, row)
	}
	return n
}

func Transpose(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func dotProduct(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func scalarMultiply(scalar float64, vec []float64) []float64 {
	result := make([]float64, len(vec))
	for i := range vec {
		result[i] = scalar * vec[i]
	}
	return result
}

func subtractVectors(a, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] - b[i]
	}
	return result
}

func normalize(vec []float64) []float64 {
	length := math.Sqrt(dotProduct(vec, vec))
	return scalarMultiply(1/length, vec)
}

func ApplyGramSchmidt(vectors [][]float64) [][]float64 {
	orthoBasis := make([][]float64, len(vectors))

	for i := range vectors {
		vec := vectors[i]
		for _, basis := range orthoBasis[:i] {
			projection := scalarMultiply(dotProduct(vec, basis), basis)
			vec = subtractVectors(vec, projection)
		}
		orthoBasis[i] = normalize(vec)
	}

	return orthoBasis
}

func ConvertIntToFloat64(data [][]int) [][]float64 {
	var n [][]float64
	for i := 0; i < len(data); i++ {
		var row []float64
		for j := 0; j < len(data[i]); j++ {
			row = append(row, float64(data[i][j]))
		}
		n = append(n, row)
	}
	return n
}

func ValidateMatrixRectangle(data [][]int) bool {
	var sizeX = len(data)
	if sizeX == 0 {
		return false
	}
	var firstSizeY int
	for i := 0; i < sizeX; i++ {
		if i == 0 {
			firstSizeY = len(data[i])
			// If the first row is empty, the matrix is invalid
			if firstSizeY == 0 {
				return false
			}
		} else if firstSizeY != len(data[i]) || len(data[i]) == 0 {
			return false
		}
	}

	return true
}

func ConvertFloat64ToFloat32(q [][]float64) [][]float32 {
	var n [][]float32
	for i := 0; i < len(q); i++ {
		var row []float32
		for j := 0; j < len(q[i]); j++ {
			row = append(row, float32(q[i][j]))
		}
		n = append(n, row)
	}
	return n
}

func Sum(a, b int) int {
	return a + b
}

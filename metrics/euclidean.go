package metrics

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// EuclideanDistance computes the euclidean distance between 2 vectors
func EuclideanDistance(x, y *mat.VecDense) float64 {
	vec := &mat.VecDense{}
	vec.AddScaledVec(x, -1, y)

	sum := 0.0

	for i := 0; i < vec.Cap(); i++ {
		sum += math.Pow(vec.AtVec(i), 2)
	}
	return sum
}

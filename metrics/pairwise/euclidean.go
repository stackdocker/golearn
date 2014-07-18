package pairwise

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

type Euclidean struct{}

func NewEuclidean() *Euclidean {
	return &Euclidean{}
}

// InnerProduct computes a Eucledian inner product.
func (self *Euclidean) InnerProduct(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	result := vectorX.Dot(vectorY)

	return result
}

// Distance computes Euclidean distance (also known as L2 distance).
func (self *Euclidean) Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	subVector := mat64.NewDense(0, 0, nil)
	subVector.Sub(vectorX, vectorY)

	result := self.InnerProduct(subVector, subVector)

	return math.Sqrt(result)
}

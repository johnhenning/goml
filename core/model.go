package core

import (
	"gonum.org/v1/gonum/mat"
)

// Model is a general interface that both Supervised and Unsupervised should conform to as well
type Model interface {
	GetWeights() ([]byte, error)
	LoadWeights(b []byte)
	Predict(X []mat.VecDense)
	PredictOne(x mat.VecDense)
}

// SupervisedModel trains on
type SupervisedModel interface {
	Fit(X []mat.VecDense, Y []mat.VecDense)
}

// UnsupervisedModel trains on unlabled data normally to cluster the data
// or to find a smaller representation of the data with dimensionality reduction
type UnsupervisedModel interface {
	Fit(X []mat.VecDense)
}

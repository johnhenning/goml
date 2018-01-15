package clustering

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/jlhbaseball15/goml/metrics"
	"gonum.org/v1/gonum/mat"
)

// KMeansClustering Model
type KMeansClustering struct {
	K              int            `json:"k"`
	Shape          int            `json:"shape"`
	DistanceMetric string         `json:"distance_metric"`
	Centroids      []mat.VecDense `json:"centroids"`
	data           []mat.VecDense
}

// NewKMeansClustering is
func NewKMeansClustering(k int, shape int, distanceMetric string) *KMeansClustering {
	centroids := make([]mat.VecDense, k)
	for i := 0; i < k; i++ {
		var centroid []float64
		for j := 0; j < shape; j++ {
			centroid = append(centroid, rand.Float64())
		}
		fmt.Println(centroid)
		centroids[i] = *mat.NewVecDense(shape, centroid)
	}
	return &KMeansClustering{
		K:              k,
		Shape:          shape,
		DistanceMetric: distanceMetric,
		Centroids:      centroids,
	}
}

// Fit takes a dataset and finds clusters
func (k *KMeansClustering) Fit(X []mat.VecDense) {
	newCentroids := newCentroids(k.K, k.Shape)
	labelCount := make([]int, k.K)
	centroidChange := 1.0
	k.data = X

	for centroidChange > 0.0 {
		centroidChange = 0.0
		// Loop over data to calculate their label
		for _, x := range X {
			centroidLabel := k.PredictOne(x)
			newCentroids[centroidLabel].AddVec(&newCentroids[centroidLabel], &x)
			labelCount[centroidLabel]++
		}

		// Calculate how much centroids moved
		for i := 0; i < k.K; i++ {
			newCentroids[i].ScaleVec(1.0/float64(labelCount[i]), &newCentroids[i])
			centroidChange += metrics.EuclideanDistance(&newCentroids[i], &k.Centroids[i])
		}

	}
}

// Predict on a whole dataset
func (k KMeansClustering) Predict(X []mat.VecDense) []int {
	labels := make([]int, len(X))
	for idx, x := range X {
		labels[idx] = k.PredictOne(x)
	}
	return labels
}

// PredictOne predicts on a single datapoint
func (k KMeansClustering) PredictOne(x mat.VecDense) int {
	minDistance := math.MaxFloat64
	index := 0
	for i, centroid := range k.Centroids {
		var currentDistance float64
		switch k.DistanceMetric {
		case "euclidean":
			currentDistance = metrics.EuclideanDistance(&x, &centroid)
		default:
			currentDistance = metrics.EuclideanDistance(&x, &centroid)
			break
		}
		if currentDistance < minDistance {
			minDistance = currentDistance
			index = i
		}
	}
	return index
}

// newCentroids is a constructor for a slice of centroids
func newCentroids(k, shape int) []mat.VecDense {
	centroids := make([]mat.VecDense, k)
	for i := 0; i < k; k++ {
		var centroid []float64
		for j := 0; j < shape; j++ {
			centroid = append(centroid, rand.Float64())
		}
		centroids[i] = *mat.NewVecDense(shape, centroid)
	}
	return centroids
}

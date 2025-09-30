package imageanalyzer

import (
	"fmt"
	"image"
	"log/slog"
	"math"
	"os"
	"runtime"
	"sync"

	"gonum.org/v1/gonum/stat"
)

var imageAnalyzerVerbose = false

func init() {
	if os.Getenv("IMAGE_ANALYZER_VERBOSE") == "true" {
		imageAnalyzerVerbose = true
	}
}

// IsGrayscale checks whether an image is grayscale to within some tolerance.
// Tolerance of 0.0 means the image must be perfectly grayscale; 0.1 allows for
// a tiny amount of color; 1.0 allows the image to be completely color.
// Returns true if the image is grayscale, false otherwise.
func IsGrayscale(img image.Image, tolerance float64) (bool, error) {
	if tolerance < 0 || tolerance > 1 {
		return false, fmt.Errorf("tolerance must be between 0 and 1, inclusive")
	}

	bounds := img.Bounds()
	totalPixels := bounds.Dx() * bounds.Dy()
	pixelRgbStdevs := make([]float64, totalPixels)

	numWorkers := runtime.GOMAXPROCS(0)
	pixelsPerWorker := totalPixels / numWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for worker := 0; worker < numWorkers; worker++ {
		startPixel := worker * pixelsPerWorker
		endPixel := startPixel + pixelsPerWorker
		if worker == numWorkers-1 {
			endPixel = totalPixels
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				x := bounds.Min.X + (i % bounds.Dx())
				y := bounds.Min.Y + (i / bounds.Dx())
				pixel := img.At(x, y)
				r, g, b, _ := pixel.RGBA()
				rgbSd := stat.StdDev([]float64{float64(r), float64(g), float64(b)}, nil)
				pixelRgbStdevs[i] = rgbSd
			}
		}(startPixel, endPixel)
	}

	wg.Wait()

	mean, _ := stat.MeanStdDev(pixelRgbStdevs, nil)
	gsScore := math.Exp(-1.0 * mean / 5000.0)

	if imageAnalyzerVerbose {
		slog.Default().Info("IsGrayscale", "gsScore", gsScore, "tolerance", tolerance, "mean", mean)
	}

	return gsScore > (1 - tolerance), nil
}

package imageanalyzer

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestIsGrayscale(t *testing.T) {
	tests := []struct {
		filename  string
		tolerance float64
		expected  bool
	}{
		{"DSCF0383.JPG", 0.1, true},
		{"DSCF0383.JPG", 0.001, false},
		{"DSCF0389.JPG", 0.2, false},
		{"DSCF0389.JPG", 1.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			path := filepath.Join("testdata", "IsGrayscale", tt.filename)
			f, err := os.Open(path)
			if err != nil {
				t.Fatalf("failed to open test image: %v", err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Fatalf("failed to decode image: %v", err)
			}

			result, err := IsGrayscale(img, tt.tolerance)
			if err != nil {
				t.Fatalf("IsGrayscale returned error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

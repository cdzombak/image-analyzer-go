# image-analyzer-go

Go library for analyzing image properties.

[![Go Reference](https://pkg.go.dev/badge/github.com/cdzombak/image-analyzer-go.svg)](https://pkg.go.dev/github.com/cdzombak/image-analyzer-go)

## Installation

```shell
go get github.com/cdzombak/image-analyzer-go
```

## Usage

### `IsGrayscale`

```go
import (
    "image"
    _ "image/jpeg"
    "os"

    imageanalyzer "github.com/cdzombak/image-analyzer-go"
)

func main() {
    f, _ := os.Open("image.jpg")
    defer f.Close()

    img, _, _ := image.Decode(f)

    isGray, _ := imageanalyzer.IsGrayscale(img, 0.1)
    if isGray {
        println("Image is grayscale")
    }
}
```

For complete documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/cdzombak/image-analyzer-go).

## License

MIT License. See [LICENSE](LICENSE).

## Author

Chris Dzombak
- https://www.dzombak.com
- https://github.com/cdzombak

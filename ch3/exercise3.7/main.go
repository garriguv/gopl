package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	flag.Parse()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + xmin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

var rootColors = []color.RGBA{
	{46, 204, 113, 255},
	{52, 152, 219, 255},
	{155, 89, 182, 255},
	{52, 73, 94, 255},
}

var selectedColor = map[complex128]color.RGBA{}

func newton(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return 4 * z * z * z
	}

	const (
		iterations = 200
	)

	z = f(z)
	for n := uint8(0); n < iterations; n++ {
		z1 := z - f(z)/fPrime(z)
		if cmplx.Abs(z1-z) < 1e-10 {
			approx := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := selectedColor[approx]
			if !ok {
				if len(rootColors) == 0 {
					log.Fatal("ran out of colors...")
				}
				c = rootColors[0]
				rootColors = rootColors[1:]
				selectedColor[approx] = c
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(n)) / math.Log(float64(iterations))
			y -= uint8(float64(y) * scale)
			return color.YCbCr{Y: y, Cb: cb, Cr: cr}
		}
		z = z1
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}

package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"os"
)

var (
	precision = flag.String("precision", "complex128", "mandelbrot precision: complex64, complex128, big.Float, or big.Rat")
	x         = flag.Float64("x", 0, "center x value")
	y         = flag.Float64("y", 0, "center y value")
	xRange    = flag.Float64("xRange", 2, "range for the x value x-range..x+range")
	yRange    = flag.Float64("yRange", 2, "range for the y value y-range..y+range")
	width     = flag.Int("width", 1024, "width of the image")
	height    = flag.Int("height", 1024, "height of the image")
)

func main() {
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	m := NewMandelbrot(*x, *y, *xRange, *yRange, *width, *height)
	switch *precision {
	case "complex64":
		m.SetFunc(mandelbrot64)
	case "complex128":
		m.SetFunc(mandelbrot128)
	case "big.Float":
		m.SetFunc(mandelbrotBigFloat)
	case "big.Rat":
		m.SetFunc(mandelbrotRat)
	default:
		log.Fatal("invalid precision: ", *precision)
	}
	m.render(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

type Mandelbrot struct {
	xmin, ymin, xmax, ymax float64
	width, height          int
	mFunc                  func(z complex128) color.Color
}

func NewMandelbrot(x, y, xRange, yRange float64, width, height int) *Mandelbrot {
	return &Mandelbrot{
		xmin:   x - xRange,
		ymin:   y - yRange,
		xmax:   x + xRange,
		ymax:   y + yRange,
		width:  width,
		height: height,
	}
}

func (m *Mandelbrot) SetFunc(f func(z complex128) color.Color) {
	m.mFunc = f
}

func (m *Mandelbrot) render(img *image.RGBA) {
	for py := 0; py < m.height; py++ {
		y := float64(py)/float64(m.height)*(m.ymax-m.ymin) + m.xmin
		for px := 0; px < m.width; px++ {
			x := float64(px)/float64(m.width)*(m.xmax-m.xmin) + m.xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, m.mFunc(z))
		}
	}
}

func mandelbrot128(z complex128) color.Color {
	const (
		iterations = 200
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return iterationColor(n, iterations)
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	const (
		iterations = 200
	)

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return iterationColor(n, iterations)
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const (
		iterations = 200
	)
	zR, zI := big.NewFloat(real(z)), big.NewFloat(imag(z))
	vR, vI := big.NewFloat(0), big.NewFloat(0)
	for n := uint8(0); n < iterations; n++ {
		vRp, vIp := big.NewFloat(0), big.NewFloat(0)
		vRp.Mul(vR, vR).Sub(vRp, (&big.Float{}).Mul(vI, vI)).Add(vRp, zR)
		vIp.Mul(vR, vI).Mul(vIp, big.NewFloat(2)).Add(vIp, zI)
		vR, vI = vRp, vIp
		hyp := big.NewFloat(0)
		hyp.Mul(vR, vR).Add(hyp, (&big.Float{}).Mul(vI, vI))
		if hyp.Cmp(big.NewFloat(4)) == 1 {
			return iterationColor(n, iterations)
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	const (
		iterations = 200
	)

	zR, zI := (&big.Rat{}).SetFloat64(real(z)), (&big.Rat{}).SetFloat64(imag(z))
	vR, vI := &big.Rat{}, &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		vRp, vIp := &big.Rat{}, &big.Rat{}
		vRp.Mul(vR, vR).Sub(vRp, (&big.Rat{}).Mul(vI, vI)).Add(vRp, zR)
		vIp.Mul(vR, vI).Mul(vIp, (&big.Rat{}).SetFloat64(2)).Add(vIp, zI)
		vR, vI = vRp, vIp
		hyp := &big.Rat{}
		hyp.Mul(vR, vR).Add(hyp, (&big.Rat{}).Mul(vI, vI))
		if hyp.Cmp((&big.Rat{}).SetFloat64(4)) == 1 {
			return iterationColor(n, iterations)
		}
	}
	return color.Black
}
func iterationColor(n uint8, iterations int) color.Color {
	factor := 255 / uint8(iterations) * n
	return color.Gray{Y: 255 - uint8(255*factor)}
}

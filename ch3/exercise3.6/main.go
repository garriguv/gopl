package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var (
	supersample = flag.Bool("supersample", false, "enable supersampling")
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
		y1 := float64(py+1)/height*(ymax-ymin) + xmin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			x1 := float64(px+1)/width*(xmax-xmin) + xmin
			za := complex(x, y)
			zb := complex(x1, y)
			zc := complex(x, y1)
			zd := complex(x, y1)
			// Image point (px, py) represents complex value z.
			if *supersample {
				colors := []color.Color{
					mandelbrot(za),
					mandelbrot(zb),
					mandelbrot(zc),
					mandelbrot(zd)}
				img.Set(px, py, avg(colors))
			} else {
				img.Set(px, py, mandelbrot(za))
			}
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := contrast * n
			return color.RGBA{R: red, G: 0, B: 0, A: 255}
		}
	}
	return color.Black
}

// https://blog.golang.org/go-image-package
func avg(colors []color.Color) color.Color {
	var r, g, b uint16
	for _, c := range colors {
		cr, cg, cb, _ := c.RGBA()
		r += uint16(cr)
		g += uint16(cg)
		b += uint16(cb)
	}
	n := uint16(len(colors))
	r /= n
	g /= n
	b /= n
	return color.RGBA{
		R: uint8(r / 0x101),
		G: uint8(g / 0x101),
		B: uint8(b / 0x101),
		A: 255,
	}
}

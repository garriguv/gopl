package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	flag.Parse()

	http.HandleFunc("/mandelbrot", mandelbrotHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	width, err := parsePosInt(r, "width")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	height, err := parsePosInt(r, "height")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	x, err := parseFloat64(r, "x")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	y, err := parseFloat64(r, "y")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	xRange, err := parseFloat64(r, "xRange")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	yRange, err := parseFloat64(r, "yRange")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := NewMandelbrot(x, y, xRange, yRange, width, height)
	m.SetFunc(mandelbrot128)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	m.render(img)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func parsePosInt(r *http.Request, name string) (value int, err error) {
	value, err = strconv.Atoi(r.FormValue(name))
	if err != nil {
		err = errors.New(fmt.Sprintf("error parsing %q: %v", name, err))
		return
	}
	if value <= 0 {
		err = errors.New(fmt.Sprintf("invalid value for %q: %d", name, value))
		return
	}
	return
}

func parseFloat64(r *http.Request, name string) (value float64, err error) {
	value, err = strconv.ParseFloat(r.FormValue(name), 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("error parsing %q: %v", name, err))
		return
	}
	return
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

func iterationColor(n uint8, iterations int) color.Color {
	factor := 255 / uint8(iterations) * n
	return color.Gray{Y: 255 - uint8(255*factor)}
}

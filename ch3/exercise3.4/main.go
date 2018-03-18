package main

import (
	"errors"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var sinAngle, cosAngle float64

func main() {
	http.HandleFunc("/surface", surfaceHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
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
	cells, err := parsePosInt(r, "cells")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	angle, err := parsePosFloat64(r, "angle")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := SurfaceParams{
		width:   width,
		height:  height,
		xyrange: 30.0,
		zbounds: 0.1,
		cells:   cells,
		angle:   angle}
	sinAngle, cosAngle = math.Sin(s.angle*math.Pi), math.Cos(s.angle*math.Pi)
	s.surface(w)
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

func parsePosFloat64(r *http.Request, name string) (value float64, err error) {
	value, err = strconv.ParseFloat(r.FormValue(name), 64)
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

type SurfaceParams struct {
	width, height int     // canvas size in pixels
	xyrange       float64 // axis ranges (-xyrange..xyrange)
	zbounds       float64 // z axis color bounds
	cells         int     // number of grid cells
	angle         float64 // angle of x, y axes (in multiples of Pi)
}

// pixels per x or y unit
func (s SurfaceParams) xyscale() float64 {
	return float64(s.width) / 2 / s.xyrange
}

// pixels per z unit
func (s SurfaceParams) zscale() float64 {
	return float64(s.height) * 0.4
}

func (s *SurfaceParams) surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", s.width, s.height)
	for i := 0; i < s.cells; i++ {
		for j := 0; j < s.cells; j++ {
			ax, ay := s.corner(i+1, j)
			bx, by := s.corner(i, j)
			cx, cy := s.corner(i, j+1)
			dx, dy := s.corner(i+1, j)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			c := s.strokeColor(i, j)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='%v'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.toHEX())
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func (s *SurfaceParams) corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(s.width)/2 + (x-y)*cosAngle*s.xyscale()
	sy := float64(s.height)/2 + (x+y)*sinAngle*s.xyscale() - z*s.zscale()
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

type StrokeColor color.RGBA

func (s *SurfaceParams) strokeColor(i, j int) StrokeColor {
	// Find point (x,y) at corner of cell (i,j).
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	return zColor(z, s.zbounds)
}

func zColor(z, zbounds float64) StrokeColor {
	if z > zbounds {
		return StrokeColor(color.RGBA{R: 255, G: 0, B: 0, A: 0})
	} else if z < -zbounds {
		return StrokeColor(color.RGBA{R: 0, G: 0, B: 255, A: 0})
	}
	r := uint8(255 - 255/(zbounds*2)*(zbounds-z))
	b := uint8(255 - 255/(zbounds*2)*(zbounds+z))
	return StrokeColor(color.RGBA{R: r, G: 0, B: b, A: 0})
}

func (c *StrokeColor) toHEX() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

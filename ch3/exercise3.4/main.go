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

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30º]
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

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

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, width, height)
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

func surface(w io.Writer, height int, width int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j, width, height)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			c := strokeColor(i, j)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='%v'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.toHEX())
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
	xyscale := width / 2 / xyrange  // pixels per x or y unit
	zscale := float64(height) * 0.4 // pixels per z unit

	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

type StrokeColor color.RGBA

func strokeColor(i, j int) StrokeColor {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	return zColor(z)
}

func zColor(z float64) StrokeColor {
	zbounds := 0.1
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

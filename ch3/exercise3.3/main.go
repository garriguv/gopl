package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30º]
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			c := strokeColor(i, j)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='%v'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.toHEX())
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
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

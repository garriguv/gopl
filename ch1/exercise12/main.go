// server1 is a minimal "echo" server
package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cycles, err := parsePosInt(r, "cycles")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	size, err := parsePosInt(r, "size")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	nframes, err := parsePosInt(r, "nframes")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	delay, err := parsePosInt(r, "delay")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	lissajous(w, cycles, size, nframes, delay)
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// lissajous generates a Lissajous GIF and writes it to w. The GIF will contain
// nframes frames separated by delays in 10ms. The size of the GIF will be
// [-size..+size]
// The number of Lissajous cycles can be adjusted using cycles.
func lissajous(out io.Writer, cycles int, size int, nframes int, delay int) {
	const (
		res = 0.001 // angular resolution
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

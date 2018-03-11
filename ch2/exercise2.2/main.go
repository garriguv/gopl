package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/garriguv/gopl/ch2/tempconv"
)

var t = flag.Bool("t", false, "convert temperatures")
var d = flag.Bool("d", false, "convert distances")

func main() {
	flag.Parse()
	if !(*t || *d) {
		fmt.Fprintln(os.Stderr, "exercise2.2: missing conversion flag(s)")
		os.Exit(1)
	}
	if len(flag.Args()) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			convert(input.Text())
		}
	} else {
		for _, v := range flag.Args() {
			convert(v)
		}
	}
}

func convert(value string) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println(os.Stderr, "exercise2.2: %v\n", err)
		os.Exit(1)
	}
	if *t {
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	if *d {
		f := Feet(v)
		m := Meters(v)
		fmt.Printf("%s = %s, %s = %s\n",
			f, FToM(f), m, MToF(m))
	}
}

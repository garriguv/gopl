package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"regexp"

	"errors"

	"github.com/garriguv/gopl/ch2/tempconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			convert(input.Text())
		}
	} else {
		for _, v := range os.Args[1:] {
			convert(v)
		}
	}
}

type measure struct {
	value float64
	unit  string
}

func convert(str string) {
	m, err := extractMeasure(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise2.2: %v\n", err)
		return
	}
	switch m.unit {
	case "C":
		c := tempconv.Celsius(m.value)
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	case "F":
		f := tempconv.Fahrenheit(m.value)
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	case "m":
		m := Meters(m.value)
		fmt.Printf("%s = %s\n", m, MToF(m))
	case "ft":
		ft := Feet(m.value)
		fmt.Printf("%s = %s\n", ft, FToM(ft))
	case "kg":
		kg := Kilogram(m.value)
		fmt.Printf("%s = %s\n", kg, KGToLBS(kg))
	case "lbs":
		lbs := Pound(m.value)
		fmt.Printf("%s = %s\n", lbs, LBSToKG(lbs))
	default:
		fmt.Fprintf(os.Stderr, "exercise2.2: unknown unit %q\n", m.unit)
	}
}

func extractMeasure(str string) (m measure, err error) {
	r, err := regexp.Compile("^(-?\\d+(?:\\.\\d+)?)([a-zA-Z]+)$")
	if err != nil {
		err = errors.New(fmt.Sprintf("exercise2.2: could not compile unit regex: %v\n", err))
		return
	}
	matches := r.FindAllStringSubmatch(str, -1)
	if matches == nil || len(matches[0]) != 3 {
		err = errors.New(fmt.Sprintf("exercise2.2: %q is not a measure\n", str))
		return
	}
	v, err := strconv.ParseFloat(matches[0][1], 64)
	if err != nil {
		return
	}
	m = measure{v, matches[0][2]}
	return
}

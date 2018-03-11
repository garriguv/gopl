package main

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }

func MToF(m Meters) Feet { return Feet(m * 3.28) }
func FToM(f Feet) Meters { return Meters(f / 3.28) }

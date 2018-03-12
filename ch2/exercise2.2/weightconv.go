package main

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string { return fmt.Sprintf("%.2fkg", kg) }
func (lbs Pound) String() string   { return fmt.Sprintf("%.2flbs", lbs) }

func KGToLBS(kg Kilogram) Pound  { return Pound(kg * 2.2) }
func LBSToKG(lbs Pound) Kilogram { return Kilogram(lbs / 2.2) }

package asciigraph

import "math"

func minFloat64Slice(v []float64) (m float64) {
	if len(v) > 0 {
		m = v[0]
	} else {
		panic("Empty slice")
	}

	for _, e := range v {
		if e < m {
			m = e
		}
	}

	return m
}

func maxFloat64Slice(v []float64) (m float64) {
	if len(v) > 0 {
		m = v[0]
	} else {
		panic("Empty slice")
	}

	for _, e := range v {
		if e > m {
			m = e
		}
	}

	return m
}

func round(input float64) float64 {
	if math.IsNaN(input) {
		return math.NaN()
	}
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}
	_, decimal := math.Modf(input)
	var rounded float64
	if decimal >= 0.5 {
		rounded = math.Ceil(input)
	} else {
		rounded = math.Floor(input)
	}
	return rounded * sign
}

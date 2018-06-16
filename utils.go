package asciigraph

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

package utils

import "strconv"

func StringsToFloats(strs []string) ([]float64, error) {
	floats := make([]float64, 0, len(strs))
	for _, s := range strs {
		if s == "" {
			continue
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}
	return floats, nil
}

func StringsToInts(strs []string) ([]int64, error) {
	ints := make([]int64, 0, len(strs))
	for _, s := range strs {
		if s == "" {
			continue
		}
		f, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		ints = append(ints, f)
	}
	return ints, nil
}

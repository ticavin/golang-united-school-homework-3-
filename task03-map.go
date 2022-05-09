package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var kk []int
	for i := range input {
		kk = append(kk, i)
	}

	sort.Ints(kk)

	for _, i := range kk {
		result = append(result, input[i])
	}
	return result
}
